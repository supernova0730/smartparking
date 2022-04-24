package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"smartparking/config"
	"smartparking/internal/models"
	"smartparking/pkg/logger"
)

var migrateCmd = &cobra.Command{
	Use:    "migrate",
	Short:  "migrate models to database",
	PreRun: loadConfig,
	Run:    migrate,
}

func migrate(cmd *cobra.Command, args []string) {
	err := config.DBConn.AutoMigrate(
		&models.Tax{},
		&models.Client{},
		&models.Car{},
		&models.ParkingZone{},
		&models.EntryHistory{},
		&models.ParkingPlace{},
		&models.Ticket{},
		&models.EmailVerification{},
		&models.Session{},
		&models.Job{},
	)
	if err != nil {
		logger.Log.Fatal("can't migrate to database", zap.Error(err))
	}

	logger.Log.Info("Models migrated")
}
