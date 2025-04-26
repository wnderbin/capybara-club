package main

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/database"
	"cap-club/admin-service/logger"
	"cap-club/admin-service/migrator"
	"cap-club/admin-service/models"
	"cap-club/admin-service/routes"
	"cap-club/admin-service/utils"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func DbInit(conf *config.ServiceConfig, log *slog.Logger) {
	sqldb, err := database.DB.DB()
	if err != nil {
		log.Error("[database] failed to get sqldb")
	}
	migrator.ApplyMigrations(sqldb)

	var adminFromConfig models.Admin

	adm_name := conf.AdminUsername
	adm_email := conf.AdminEmail
	adm_password := conf.AdminPassword
	hashed_password, err := utils.HashPassword(adm_password)
	if err != nil {
		log.Error(err.Error())
	}
	err = database.DB.Where("name = ?", adm_name).First(&adminFromConfig).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(err.Error())
	} else if err == gorm.ErrRecordNotFound {
		database.DB.Create(&models.Admin{
			Id:       uuid.NewString(),
			Name:     adm_name,
			Email:    adm_email,
			Password: hashed_password,
		})
	} else {
		log.Info("Admin already exists")
	}
}

func main() {
	conf := config.MustLoad()
	log := logger.LoggerInit(conf.Env)
	log = log.With(slog.String("env", conf.Env))
	if conf.StartUpStatus == 0 {
		log.Info("+")
	} else {
		DbInit(conf, log)
		log.Info("Initializing service", slog.String("Address", fmt.Sprintf("%s:%d", conf.Address, conf.Port)))
		router := gin.Default()
		router.LoadHTMLGlob("admin-service/ui/html/*")
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
