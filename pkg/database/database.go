package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"smartparking/pkg/logger"
)

func Init(conf Config) (*gorm.DB, error) {
	dsn := conf.makeDSN()
	gormDialector := postgres.New(postgres.Config{DSN: dsn})
	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         logger.GormLogger,
	}
	return gorm.Open(gormDialector, gormConfig)
}
