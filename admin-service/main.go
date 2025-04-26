package main

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/database"
	"cap-club/admin-service/logger"
	"cap-club/admin-service/migrator"
	"cap-club/admin-service/routes"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
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
		router := gin.Default()
		router.LoadHTMLGlob("admin-service/ui/html/*")
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
