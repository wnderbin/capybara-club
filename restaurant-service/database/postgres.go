package database

import (
	"cap-club/restaurant-service/config"
	"cap-club/restaurant-service/logger"
	"fmt"
	"log/slog"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	conf *config.ServiceConfig = config.MustLoad()
	Log  slog.Logger           = *logger.LoggerInit(conf.Env)
	DB   *gorm.DB              = openPostgres()
)

func openPostgres() *gorm.DB {
	db, err := initPostgres()
	if err != nil {
		Log.Error(err.Error())
		return nil
	}
	return db
}

func ClosePostgres() error {
	sqldb, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from gorm.DB: %w", err)
	}
	Log.Info("[postgres] closing connection with Postgres...")
	return sqldb.Close()
}

func initPostgres() (*gorm.DB, error) {
	if conf.StartUpStatus == 0 {
		return nil, nil
	}
	dsn := conf.Postgres.GetDSN()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("[postgres] failed to connect to postgres: %w", err)
	}
	sqldb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("[postgres] failed to connect to postgres: %w", err)
	}
	if err = sqldb.Ping(); err != nil {
		return nil, fmt.Errorf("[postgres] failed to ping database: %w", err)
	}

	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(5)
	sqldb.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
