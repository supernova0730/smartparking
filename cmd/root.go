package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"smartparking/config"
	"smartparking/pkg/cache"
	"smartparking/pkg/database"
	"smartparking/pkg/logger"
)

var rootCmd = &cobra.Command{
	Use:   "smartparking",
	Short: "smartparking applications",
}

func init() {
	rootCmd.AddCommand(
		startCmd,
		migrateCmd,
		configCmd,
	)

	rootCmd.PersistentFlags().StringVar(&config.GlobalConfig.ConfigPath, "config", ConfigPath, "path to config file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logger.Log.Error("executing root command", zap.Error(err))
	}
}

func loadConfig(cmd *cobra.Command, args []string) {
	configPath := config.GlobalConfig.ConfigPath
	err := config.Load(configPath)
	if err != nil {
		logger.Log.Fatal("can't initialize config", zap.Error(err))
	}

	logger.Log.Info("root.loadConfig", zap.String("config path", configPath), zap.Any("config", config.GlobalConfig))

	db, err := database.Conn(database.Config{
		Host:     config.GlobalConfig.DB.Host,
		Port:     config.GlobalConfig.DB.Port,
		Name:     config.GlobalConfig.DB.Name,
		Username: config.GlobalConfig.DB.Username,
		Password: config.GlobalConfig.DB.Password,
		SSLMode:  config.GlobalConfig.DB.SSLMode,
	})
	if err != nil {
		logger.Log.Fatal("can't initialize database connection", zap.Error(err))
	}

	mcHost := config.GlobalConfig.Cache.Host
	mcPort := config.GlobalConfig.Cache.Port
	mc, err := cache.Conn(mcHost, mcPort)
	if err != nil {
		logger.Log.Fatal("can't initialize cache connection", zap.Error(err))
	}

	config.DBConn = db
	config.CacheConn = mc

	logger.Log.Info("Database connection initialized...")
	logger.Log.Info("Cache connection initialized...")
}
