package commands

import (
	"context"
	"fmt"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli/httpcfg"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"github.com/ledgerwatch/erigon/rpc"
	v3log "github.com/ledgerwatch/log/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"runtime/trace"
)

var ExportCmd = &cobra.Command{
	Use: "export",
}

func init() {
}

func GetExportCmd(cfg *httpcfg.HttpCfg, rootCancel context.CancelFunc) *cobra.Command {
	var (
		outJsonL    bool
		outProtobuf bool
		outputDir   string
	)

	var (
		exportTrace string
	)

	//	exportTrace := &cobra.Command{
	//		Use: "trace [start] [end(optinal)]",
	//		Example: `./rpcdaemon export trace 122 ( export trace only on height: 122)
	//./rpcdaemon export trace 122 222 ( export trace from height: 122 to height: 222)`,
	//		Args: cobra.RangeArgs(1, 2),
	//		Run: func(cmd *cobra.Command, args []string) {
	//			logger := log.New()
	//			var startBlock, endBlock rpc.BlockNumber
	//			_ = endBlock
	//			if err := startBlock.UnmarshalJSON([]byte(args[0])); err != nil {
	//
	//				logger.Error("parse startBlock failed", err)
	//				return
	//			}
	//			if len(args) == 2 {
	//				logger.Error("export trace from height: ", args[0], " to height: ", args[1], " is not support now")
	//				return
	//			}
	//			ctx := cmd.Context()
	//			db, borDb, _, _, _, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *cfg, logger, rootCancel)
	//			if err != nil {
	//				log.Error("Could not connect to DB", "error", err)
	//				os.Exit(1)
	//			}
	//			defer db.Close()
	//			if borDb != nil {
	//				defer borDb.Close()
	//			}
	//
	//			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
	//			if cfg.TevmEnabled {
	//				base.EnableTevmExperiment()
	//			}
	//
	//			var file *os.File
	//			if outputFile != "" {
	//				file, err = os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644)
	//				if err != nil {
	//					logger.Error("open file failed", "error", err)
	//					return
	//				}
	//			}
	//			defer file.Close()
	//
	//			stream := jsoniter.NewStream(jsoniter.ConfigDefault, file, 0)
	//			defer stream.Flush()
	//
	//			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)
	//			tracerName := "opsTracer"
	//			ctx = context.WithValue(ctx, "logger", logger)
	//			if err := debugImpl.TraceSingleBlock(ctx, startBlock, &tracers.TraceConfig{
	//				Tracer: &tracerName,
	//			}, stream); err != nil {
	//				logger.Error("Could not trace block", "error", err)
	//			} else {
	//				logger.Info("trace block success")
	//			}
	//		},
	//	}
	//
	//	exportBlock := &cobra.Command{
	//		Use: "block [start] [end(optional)]",
	//		Example: `./rpcdaemon export block 122 ( export block only on height: 122)
	//./rpcdaemon export block 122 222 ( export block from height: 122 to height: 222)`,
	//		Args: cobra.RangeArgs(1, 2),
	//		Run: func(cmd *cobra.Command, args []string) {
	//			logger := log.New()
	//			ctx := cmd.Context()
	//			db, borDb, eth, txPool, mining, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *cfg, logger, rootCancel)
	//			if err != nil {
	//				log.Error("Could not connect to DB", "error", err)
	//				os.Exit(1)
	//			}
	//			defer db.Close()
	//			if borDb != nil {
	//				defer borDb.Close()
	//			}
	//
	//			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
	//			if cfg.TevmEnabled {
	//				base.EnableTevmExperiment()
	//			}
	//			ethImpl := NewEthAPI(base, db, eth, txPool, mining, cfg.Gascap)
	//
	//			block, err := ethImpl.GetFullBlockByNumber(ctx, 0, true)
	//			if err != nil {
	//				logger.Error("Could not get block number", "error", err)
	//				return
	//			}
	//			logger.Info(block.Hash)
	//
	//		},
	//	}
	//
	zlog := zap.L().Sugar()
	exportTx := &cobra.Command{
		Use: "tx [start] [end(optional)]",
		Example: `./rpcdaemon export tx 122 ( export tx only on height: 122)
./rpcdaemon export block 122 222 ( export tx from height: 122 to height: 222)`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			if exportTrace != "" {
				runtimeTrace, err := os.OpenFile(exportTrace, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
				if err != nil {
					zlog.Errorf("open file %s %s", exportTrace, err)
				}
				if err := trace.Start(runtimeTrace); err != nil {
					zlog.Errorf("start trace %s", err)
				}
				defer trace.Stop()
			}
		
			logger := v3log.New()

			var startBlock, endBlock rpc.BlockNumber
			_ = endBlock
			if err := startBlock.UnmarshalJSON([]byte(args[0])); err != nil {
				return
			}
			endBlock = startBlock
			if len(args) == 2 {
				if err := endBlock.UnmarshalJSON([]byte(args[1])); err != nil {
					zlog.Errorf("parse endBlock failed %s", err)
					return
				}
			}

			ctx := cmd.Context()
			db, borDb, eth, txPool, mining, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *cfg, logger, rootCancel)
			if err != nil {
				zlog.Errorf("Could not connect to DB: %s", err)
				os.Exit(1)
			}
			_ = eth
			_ = txPool
			_ = mining
			defer db.Close()
			if borDb != nil {
				defer borDb.Close()
			}

			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
			if cfg.TevmEnabled {
				base.EnableTevmExperiment()
			}
			//ethImpl := NewEthAPI(base, db, eth, txPool, mining, cfg.Gascap)

			filename := fmt.Sprintf("%d", startBlock)
			if endBlock != 0 {
				filename = fmt.Sprintf("%d-%d", startBlock, endBlock)
			}

			if outProtobuf {
				filename += ".proto.bin"
			} else {
				filename += ".jsonl"
			}

			path := filepath.Join(outputDir, filename)
			file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				zlog.Errorf("open file failed: %s", err)
				return
			}
			defer file.Close()
			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)

			tracerName := "opsTracer"
			ctx = context.WithValue(ctx, "logger", logger)
			for height := startBlock; height <= endBlock; height++ {
				if err := debugImpl.TraceSingleBlock(ctx, startBlock, &tracers.TraceConfig{
					Tracer: &tracerName,
				}, file, outProtobuf); err != nil {
					zlog.Errorf("could not trace block: %s", err)

				} else {
					zlog.Infof("trace block %d success", height)
				}
			}
			zlog.Info("export transactions data to ", path)
		},
	}
	//exportTrace.PersistentFlags().StringVar(&outputFile, "output", "", "output file to save export data")
	//exportBlock.PersistentFlags().StringVar(&outputFile, "output", "", "output file to save export data")
	exportTx.PersistentFlags().BoolVar(&outJsonL, "out-jsonl", true, "save export data as jsonl file")
	exportTx.PersistentFlags().BoolVar(&outProtobuf, "out-proto", false, "save export data as protobuf serialized file")

	exportTx.PersistentFlags().StringVar(&outputDir, "out-dir", ".", "save export data to a dir ")

	//xExportCmd.AddCommand(exportBlock)
	//xExportCmd.AddCommand(exportTrace)
	ExportCmd.AddCommand(exportTx)
	ExportCmd.PersistentFlags().StringVar(&exportTrace, "runtime-trace", "", "golang process runtime trace")

	return ExportCmd
}
