package database

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/logger"
	"cap-club/admin-service/models"
	"cap-club/admin-service/utils"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
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
	var adminFromConfig models.Admin
	adm_name := conf.AdminUsername
	adm_email := conf.AdminEmail
	adm_password := conf.AdminPassword
	hashed_password, err := utils.HashPassword(adm_password)
	if err != nil {
		Log.Error(err.Error())
		return nil
	}
	err = db.Where("name = ?", adm_name).First(&adminFromConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Log.Error(err.Error())
		return nil
	} else if err == gorm.ErrRecordNotFound {
		db.Create(&models.Admin{
			Id:       uuid.NewString(),
			Name:     adm_name,
			Email:    adm_email,
			Password: hashed_password,
		})
	} else {
		Log.Info("Admin already exists")
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
