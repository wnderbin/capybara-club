package main

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/routes"
	"cap-club/database"
	"cap-club/migrator"
	"cap-club/models"
	"cap-club/utils"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	conf := config.MustLoad()
	if database.Conf.StartUpStatus == 0 {
		database.Log.Info("+")
	} else {
		sqldb, err := database.DB.DB()
		if err != nil {
			database.Log.Error("[database] failed to get sqldb")
		}
		migrator.ApplyMigrations(sqldb)
		database.Log.Info("Initializing service", slog.String("Address", fmt.Sprintf("%s:%d", conf.Address, conf.Port)))
		hashed_password, err := utils.HashPassword(conf.AdminPassword)
		if err != nil {
			database.Log.Error("[utils] cannot hash password")
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
