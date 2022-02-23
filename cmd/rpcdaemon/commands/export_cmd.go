package commands

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli/httpcfg"
	"github.com/ledgerwatch/erigon/eth/tracers"
	"github.com/ledgerwatch/erigon/rpc"
	"github.com/ledgerwatch/log/v3"
	"github.com/spf13/cobra"
	"os"
)

var xExportCmd = &cobra.Command{
	Use: "export",
}

func init() {
}

func ExportCmd(cfg *httpcfg.HttpCfg, rootCancel context.CancelFunc) *cobra.Command {
	var outputFile string

	exportTrace := &cobra.Command{
		Use: "trace [start] [end(optinal)]",
		Example: `./rpcdaemon export trace 122 ( export trace only on height: 122)
./rpcdaemon export trace 122 222 ( export trace from height: 122 to height: 222)`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.New()
			var startBlock, endBlock rpc.BlockNumber
			_ = endBlock
			if err := startBlock.UnmarshalJSON([]byte(args[0])); err != nil {

				logger.Error("parse startBlock failed", err)
				return
			}
			if len(args) == 2 {
				logger.Error("export trace from height: ", args[0], " to height: ", args[1], " is not support now")
				return
			}
			ctx := cmd.Context()
			db, borDb, _, _, _, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *cfg, logger, rootCancel)
			if err != nil {
				log.Error("Could not connect to DB", "error", err)
				os.Exit(1)
			}
			defer db.Close()
			if borDb != nil {
				defer borDb.Close()
			}

			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
			if cfg.TevmEnabled {
				base.EnableTevmExperiment()
			}

			var file *os.File
			if outputFile != "" {
				file, err = os.OpenFile(outputFile, os.O_CREATE|os.O_RDWR, 0644)
				if err != nil {
					logger.Error("open file failed", "error", err)
					return
				}
			}
			defer file.Close()

			stream := jsoniter.NewStream(jsoniter.ConfigDefault, file, 0)
			defer stream.Flush()

			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)
			tracerName := "opsTracer"
			ctx = context.WithValue(ctx, "logger", logger)
			if err := debugImpl.TraceSingleBlock(ctx, startBlock, &tracers.TraceConfig{
				Tracer: &tracerName,
			}, stream); err != nil {
				logger.Error("Could not trace block", "error", err)
			} else {
				logger.Info("trace block success")
			}
		},
	}

	exportBlock := &cobra.Command{
		Use: "block [start] [end(optional)]",
		Example: `./rpcdaemon export block 122 ( export block only on height: 122)
./rpcdaemon export block 122 222 ( export block from height: 122 to height: 222)`,
		Args: cobra.RangeArgs(1, 2),
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.New()
			ctx := cmd.Context()
			db, borDb, eth, txPool, mining, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *cfg, logger, rootCancel)
			if err != nil {
				log.Error("Could not connect to DB", "error", err)
				os.Exit(1)
			}
			defer db.Close()
			if borDb != nil {
				defer borDb.Close()
			}

			base := NewBaseApi(filters, stateCache, blockReader, cfg.SingleNodeMode)
			if cfg.TevmEnabled {
				base.EnableTevmExperiment()
			}
			ethImpl := NewEthAPI(base, db, eth, txPool, mining, cfg.Gascap)

			block, err := ethImpl.GetBlockByNumber(ctx, 0, true)
			if err != nil {
				logger.Error("Could not get block number", "error", err)
				return
			}
			for k, v := range block {
				logger.Info("block", "key", k, "value", v)
			}
		},
	}
	exportTrace.PersistentFlags().StringVar(&outputFile, "output", "", "output file to save export data")
	exportBlock.PersistentFlags().StringVar(&outputFile, "output", "", "output file to save export data")

	xExportCmd.AddCommand(exportBlock)
	xExportCmd.AddCommand(exportTrace)

	return xExportCmd
}
