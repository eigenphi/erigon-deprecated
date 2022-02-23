package commands

import (
	"context"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli/httpcfg"
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
	exportTrace := &cobra.Command{
		Use: "trace",
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.New()
			ctx := cmd.Context()
			db, borDb, _, _, _, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *httpcfg.Cfg, logger, rootCancel)
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

			debugImpl := NewPrivateDebugAPI(base, db, cfg.Gascap)
			debugImpl.TraceBlockByRange(ctx, 0, 1, nil, nil)
		},
	}

	exportBlock := &cobra.Command{
		Use:  "block [startBlock] [endBlock]",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			logger := log.New()
			ctx := cmd.Context()
			db, borDb, eth, txPool, mining, _, stateCache, blockReader, filters, err := cli.RemoteServices(ctx, *httpcfg.Cfg, logger, rootCancel)
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

			ethImpl.GetBlockByNumber(ctx, 0, true)
		},
	}

	xExportCmd.AddCommand(exportBlock)
	xExportCmd.AddCommand(exportTrace)

	return xExportCmd
}
