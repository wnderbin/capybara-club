package main

import (
	"cap-club/migrator"
	"cap-club/restaurant-service/config"
	"cap-club/restaurant-service/database"
	"cap-club/restaurant-service/handlers"
	"cap-club/restaurant-service/logger"
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
		go handlers.UpdateRestaurants()
		router := gin.Default()
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
