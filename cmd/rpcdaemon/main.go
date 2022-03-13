package main

import (
	"fmt"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/cli"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/commands"
	"github.com/ledgerwatch/erigon/cmd/rpcdaemon/filters"
	"github.com/ledgerwatch/erigon/cmd/utils"
	"github.com/ledgerwatch/erigon/params"
	"github.com/ledgerwatch/log/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)

func init() {
	log, err := zap.NewDevelopment(zap.AddStacktrace(zap.PanicLevel))
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(log)
}

func main() {
	cmd, cfg := cli.RootCommand()
	rootCtx, rootCancel := utils.RootContext()
	cmd.AddCommand(commands.GetExportCmd(cfg, rootCancel))
	cmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Print the version of rpcdaemon",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("rpcdaemon", "branch", params.GitBranch,
				"tag", params.GitTag, "commit", params.GitCommit)
			os.Exit(0)
		},
	})
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		logger := log.New()
		db, backend, txPool, mining, stateCache, err := cli.RemoteServices(cmd.Context(), *cfg, logger, rootCancel)
		if err != nil {
			log.Error("Could not connect to DB", "error", err)
			return nil
		}
		defer db.Close()

		var ff *filters.Filters
		if backend != nil {
			ff = filters.New(rootCtx, backend, txPool, mining)
		} else {
			log.Info("filters are not supported in chaindata mode")
		}

		if err := cli.StartRpcServer(cmd.Context(), *cfg, commands.APIList(cmd.Context(), db, backend, txPool, mining, ff, stateCache, *cfg, nil)); err != nil {
			log.Error(err.Error())
			return nil
		}
		return nil
	}

	if err := cmd.ExecuteContext(rootCtx); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
