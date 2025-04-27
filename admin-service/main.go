package main

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/database"
	"cap-club/admin-service/logger"
	"cap-club/admin-service/models"
	"cap-club/admin-service/routes"
	"cap-club/admin-service/utils"
	"cap-club/migrator"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	conf := config.MustLoad()
	log := logger.LoggerInit(conf.Env)
	log = log.With(slog.String("env", conf.Env))
	if conf.StartUpStatus == 0 {
		log.Info("+")
	} else {
		sqldb, err := database.DB.DB()
		if err != nil {
			log.Error("[database] failed to get sqldb")
		}
		migrator.ApplyMigrations(sqldb)
		log.Info("Initializing service", slog.String("Address", fmt.Sprintf("%s:%d", conf.Address, conf.Port)))
		hashed_password, err := utils.HashPassword(conf.AdminPassword)
		if err != nil {
			log.Error("[utils] cannot hash password")
		}
		database.DB.Create(&models.Admin{
			Id:       uuid.NewString(),
			Name:     conf.AdminUsername,
			Email:    conf.AdminEmail,
			Password: hashed_password,
		})
		router := gin.Default()
		router.LoadHTMLGlob("admin-service/ui/html/*")
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
