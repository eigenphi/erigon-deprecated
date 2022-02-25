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
	zlog := zap.L().Sugar()

	exportBlock := &cobra.Command{
		Use: "block [start] [end(optional)]",
		Example: `./rpcdaemon export block 122 ( export transaction only on height: 122)
./rpcdaemon export block 122 222 ( export transaction from height: 122 to height: 222)`,
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
			_, _, _ = eth, txPool, mining
			defer db.Close()
			if borDb != nil {
				defer borDb.Close()
			}

			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
			if cfg.TevmEnabled {
				base.EnableTevmExperiment()
			}
			//ethImpl := NewEthAPI(base, db, eth, txPool, mining, cfg.Gascap)

			filename := fmt.Sprintf("block_%d", startBlock)
			if endBlock != 0 {
				filename = fmt.Sprintf("block_%d-%d", startBlock, endBlock)
			}

			if outProtobuf {
				filename += ".proto.bin"
			} else {
				filename += ".jsonl"
			}

			path := filepath.Join(outputDir, filename)
			file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				zlog.Errorf("open file failed: %s", err)
				return
			}
			defer file.Close()
			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)

			for height := startBlock; height <= endBlock; height++ {
				if err := debugImpl.getSimpleBlock(ctx, height, file); err != nil {
					zlog.Errorf("could not get simple block on height: %d, %s", height, err)

				} else {
					zlog.Infof("export simple block on %d success", height)
				}
			}
			zlog.Info("export simple transactions data to ", path)
		},
	}

	exportTx := &cobra.Command{
		Use: "simple_transaction [start] [end(optional)]",
		Example: `./rpcdaemon export simple_transaction 122 ( export transaction only on height: 122)
./rpcdaemon export simple_transaction 122 222 ( export transaction from height: 122 to height: 222)`,
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
			_, _, _ = eth, txPool, mining
			defer db.Close()
			if borDb != nil {
				defer borDb.Close()
			}

			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
			if cfg.TevmEnabled {
				base.EnableTevmExperiment()
			}
			//ethImpl := NewEthAPI(base, db, eth, txPool, mining, cfg.Gascap)

			filename := fmt.Sprintf("simple_transaction_%d", startBlock)
			if endBlock != 0 {
				filename = fmt.Sprintf("simple_transaction_%d-%d", startBlock, endBlock)
			}

			if outProtobuf {
				filename += ".proto.bin"
			} else {
				filename += ".jsonl"
			}

			path := filepath.Join(outputDir, filename)
			file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				zlog.Errorf("open file failed: %s", err)
				return
			}
			defer file.Close()
			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)

			for height := startBlock; height <= endBlock; height++ {
				if err := debugImpl.getSimpleTx(ctx, height, file); err != nil {
					zlog.Errorf("could not get simple transaction on height: %d, %s", height, err)

				} else {
					zlog.Infof("export simple transaction on %d success", height)
				}
			}
			zlog.Info("export simple transactions data to ", path)
		},
	}

	exportTxTrace := &cobra.Command{
		Use: "transaction_trace [start] [end(optional)]",
		Example: `./rpcdaemon export transaction_trace 122 ( export transaction_trace only on height: 122)
./rpcdaemon export transaction_trace 122 222 ( export transaction_trace from height: 122 to height: 222)`,
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

			filename := fmt.Sprintf("transaction_trace_%d", startBlock)
			if endBlock != 0 {
				filename = fmt.Sprintf("transaction_trace_%d-%d", startBlock, endBlock)
			}

			if outProtobuf {
				filename += ".proto.bin"
			} else {
				filename += ".jsonl"
			}

			path := filepath.Join(outputDir, filename)
			file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0644)
			if err != nil {
				zlog.Errorf("open file failed: %s", err)
				return
			}
			defer file.Close()
			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)

			tracerName := "opsTracer"
			ctx = context.WithValue(ctx, "logger", logger)
			for height := startBlock; height <= endBlock; height++ {
				if err := debugImpl.TraceSingleBlock(ctx, height, &tracers.TraceConfig{
					Tracer: &tracerName,
				}, file, outProtobuf); err != nil {
					zlog.Errorf("could not trace block: %s", err)

				} else {
					zlog.Infof("trace block %d success", height)
				}
			}
			zlog.Info("export transactions trace data to ", path)
		},
	}
	exportTxTrace.PersistentFlags().BoolVar(&outJsonL, "out-jsonl", true, "save export data as jsonl file")
	exportTxTrace.PersistentFlags().BoolVar(&outProtobuf, "out-proto", false, "save export data as protobuf serialized file")

	exportBlock.PersistentFlags().StringVar(&outputDir, "out-dir", ".", "save export data to a dir ")
	exportTx.PersistentFlags().StringVar(&outputDir, "out-dir", ".", "save export data to a dir ")
	exportTxTrace.PersistentFlags().StringVar(&outputDir, "out-dir", ".", "save export data to a dir ")

	ExportCmd.AddCommand(exportBlock)
	ExportCmd.AddCommand(exportTx)
	ExportCmd.AddCommand(exportTxTrace)
	ExportCmd.PersistentFlags().StringVar(&exportTrace, "runtime-trace", "", "golang process runtime trace")

	return ExportCmd
}