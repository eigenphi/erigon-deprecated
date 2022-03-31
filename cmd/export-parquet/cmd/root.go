/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"io/fs"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	log *zap.SugaredLogger
)

const (
	ParquetDataDirFlag = "parquet-data-dir"
	ErigonDataDirFlag  = "erigon-data-dir"
	RpcDaemonCliFlag   = "rpcdaemon-cli"
)

var (
	ErigonParquetDataDir       string
	ErigonRecentParquetDataDir string

	ErigonDataDir    string
	RpcdaemonBinPath string

	MinExportInterval int64

	SafeBlockGap int64
)

func erigonLatestHeight() (int64, error) {
	buffer := bytes.NewBuffer(make([]byte, 0))
	if err := json.NewEncoder(buffer).Encode(JsonrpcParam{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "eth_blockNumber",
	}); err != nil {
		return 0, fmt.Errorf("json.NewEncoder: %v", err)
	}

	resp, err := http.Post("http://127.0.0.1:8845", "application/json", buffer)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result jsonrpcResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, fmt.Errorf("json.NewDecoder: %v", err)
	}
	return result.Result.Int64(), nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "export-parquet",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		ErigonRecentParquetDataDir = filepath.Join(ErigonParquetDataDir, "recent")
		if s, err := os.Stat(ErigonRecentParquetDataDir); err != nil {
			log.Errorf("recentDataDir not exist")
			os.Exit(1)
		} else if !s.IsDir() {
			log.Errorf("%s is not a dir", ErigonRecentParquetDataDir)
			os.Exit(1)
		}

		log.Infof("ErigonParquetDataDir: %s", ErigonParquetDataDir)
		log.Infof("ErigonRecentParuqetDataDir: %s", ErigonRecentParquetDataDir)
		log.Infof("SafeBlockGap: %d", SafeBlockGap)
		log.Infof("MinExportInterval: %d", MinExportInterval)
		log.Infof("RpcdaemonBinPath: %s", RpcdaemonBinPath)

		stat, err := os.Stat(ErigonParquetDataDir)
		if err != nil {
			return err
		}
		if !stat.IsDir() {
			return fmt.Errorf("%s is not a directory", ErigonParquetDataDir)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		latestHeight, err := erigonLatestHeight()
		if err != nil {
			return fmt.Errorf("erigonLatestHeight: %v", err)
		}
		fmt.Println("latestHeight:", latestHeight)

		var (
			archiveEnd = uint64(0)
			recentEnd  = uint64(0)
		)

		{
			var brs = make(blockRanges, 0)
			archivePaths := getParquetFilePaths(ErigonParquetDataDir)
			brs.fromPaths(archivePaths)
			archiveEnd = brs[len(brs)-1].end
		}
		{
			var brs = make(blockRanges, 0)
			recentPaths := getParquetFilePaths(ErigonRecentParquetDataDir)
			brs.fromPaths(recentPaths)
			if len(brs) == 0 {
				recentEnd = archiveEnd
			} else {
				recentEnd = brs[len(brs)-1].end
			}
		}

		return runUpdate(int64(archiveEnd), int64(recentEnd))
	},
}

var (
	ExistCh = make(chan os.Signal, 1)
)

func runUpdate(archiveEnd, recentEnd int64) error {
	log.Infof("run update archiveEnd: %d, recentEnd: %d", archiveEnd, recentEnd)
	ticker := time.NewTicker(time.Second * 10)
	lastEndBlock := recentEnd

	for {
		select {
		case <-ticker.C:
			erigonLatestHeight, err := erigonLatestHeight()
			if err != nil {
				log.Errorf("erigonLatestHeight: %v", err)
				continue
			}
			log.Infof("erigonLatestHeight %d", erigonLatestHeight)
			startBlock := lastEndBlock + 1
			endBlock := erigonLatestHeight - SafeBlockGap
			if endBlock < startBlock {
				continue
			}
			if endBlock > archiveEnd+10000 {
				log.Infof("need export archive [%d,%d]", archiveEnd+1, archiveEnd+10000)
				endBlock = archiveEnd + 10000
				if err := executeExport(ErigonParquetDataDir, archiveEnd+1, archiveEnd+10000); err != nil {
					log.Errorf("execute archive export %s", err)
				} else {
					archiveEnd = endBlock
					lastEndBlock = endBlock
				}
			} else {
				if endBlock-startBlock+1 < MinExportInterval {
					continue
				}
				if err := executeExport(ErigonRecentParquetDataDir, startBlock, endBlock); err != nil {
					log.Errorf("execte parquet export %s", err)
					continue
				}
				lastEndBlock = endBlock
			}
		case ch := <-ExistCh:
			return fmt.Errorf("runUpdate received a signal: %v", ch.String())
		}
	}
}

type blockRanges []blockRange

func (brs *blockRanges) fromPaths(path []string) {
	for _, path := range append(path) {
		br := blockRange{}
		if nscanf, err := fmt.Sscanf(path, "trace_parquet_%d-%d.parquet", &br.start, &br.end); err != nil || nscanf != 2 {
			log.Warnf("path: %s, not parse properly", path)
			continue
		} else if br.start > br.end {
			log.Warnf("path: %s, start: %d > end: %d", path, br.start, br.end)
			continue
		}
		*brs = append(*brs, br)
	}
	sort.Slice(*brs, func(i, j int) bool {
		if (*brs)[i].start == (*brs)[j].start {
			return (*brs)[i].end < (*brs)[j].end
		}
		return (*brs)[i].start < (*brs)[j].start
	})
}

type blockRange struct {
	start uint64
	end   uint64
}

func executeExport(dataStore string, start, end int64) error {
	log.Infof("executeExport: %d, %d to %s", start, end, dataStore)
	//./erigon-rpcdaemon-ethereum.v1.3.0 --datadir /mnt/data/erigon-ethereum export parquet 14460001 14470000 --out-dir /mnt/data/erigon-parquet/
	args := []string{RpcdaemonBinPath, "--datadir", ErigonDataDir, "export", "parquet", fmt.Sprint(start), fmt.Sprint(end), "--out-dir", dataStore}
	log.Infof("execute: %s", strings.Join(args, " "))
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func getParquetFilePaths(datadir string) []string {
	paths := make([]string, 0)
	filepath.WalkDir(datadir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}
		base := filepath.Base(path)
		if strings.HasPrefix(base, "trace_parquet_") && strings.HasSuffix(base, ".parquet") {
			paths = append(paths, base)
		}
		return nil
	})
	return paths
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().StringVar(&ErigonParquetDataDir, ParquetDataDirFlag, "", "parquet data dir")
	if err := rootCmd.MarkPersistentFlagRequired(ParquetDataDirFlag); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().StringVar(&ErigonDataDir, ErigonDataDirFlag, "", "erigon data dir")
	if err := rootCmd.MarkPersistentFlagRequired(ErigonDataDirFlag); err != nil {
		panic(err)
	}
	rootCmd.PersistentFlags().Int64Var(&MinExportInterval, "min-export-interval", 7, "min export interval")

	rootCmd.PersistentFlags().StringVar(&RpcdaemonBinPath, RpcDaemonCliFlag, "", "rpcdaemon cli")
	if err := rootCmd.MarkPersistentFlagRequired(RpcDaemonCliFlag); err != nil {
		panic(err)
	}

	rootCmd.PersistentFlags().Int64Var(&SafeBlockGap, "safe-block-gap", 100, "safe block gap")

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	_l, err := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		panic(err)
	}
	log = _l.Sugar()

	signal.Notify(ExistCh, os.Interrupt, syscall.SIGTERM)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.export-parquet.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
