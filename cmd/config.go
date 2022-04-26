package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"smartparking/config"
	"smartparking/pkg/logger"
	"smartparking/pkg/tools"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "initial config generation",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := cmd.Flags().GetString("path")
		if err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		if err = viper.Unmarshal(&config.GlobalConfig); err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		err = config.SaveConfig(configPath)
		if err != nil {
			logger.Log.Fatal("generate config", zap.Error(err))
		}

		logger.Log.Info("config file generated")
	},
}

func init() {
	viper.SetEnvPrefix("sp")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	confCmd := configCmd.PersistentFlags()

	confCmd.String("path", tools.FilepathFromHome(".smartparking/config.test.yml"), "generate config file to (default .smartparking/config.yml)")
}
