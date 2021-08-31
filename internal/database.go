package internal

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	log "github.com/sirupsen/logrus"
)

func NewDBConnection(host string, port string, user string, password string, DBName string, dbLogger *log.Logger) (*gorm.DB, error) {
	gormLogger := logger.New(
		dbLogger,
		logger.Config{
			LogLevel: logger.Silent,
		},
	)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: gormLogger})
	if err != nil {
		return db, fmt.Errorf("unable to connect to datbase: %s", err.Error())
	}

	dbInstance, err := db.DB()
	if err != nil {
		return db, fmt.Errorf("unable to retrieve db instance: %s", err.Error())
	}

	dbInstance.SetMaxOpenConns(10)
	dbInstance.SetMaxIdleConns(10)
	dbInstance.SetConnMaxLifetime(time.Hour)

	return db, nil
}
