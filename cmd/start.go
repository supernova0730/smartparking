package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"smartparking/internal/jobs"
	"smartparking/internal/transport"
	"smartparking/pkg/logger"
)

var startCmd = &cobra.Command{
	Use:    "start",
	Short:  "start server",
	PreRun: loadConfig,
	Run: func(cmd *cobra.Command, args []string) {
		err := jobs.Start()
		if err != nil {
			logger.Log.Fatal("can't start jobs", zap.Error(err))
		}

		err = transport.ServerStart()
		if err != nil {
			logger.Log.Fatal("can't start server", zap.Error(err))
		}
	},
}
