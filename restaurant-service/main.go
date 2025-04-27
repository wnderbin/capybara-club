package main

import (
	"cap-club/migrator"
	"cap-club/restaurant-service/config"
	"cap-club/restaurant-service/database"
	"cap-club/restaurant-service/logger"
	"cap-club/restaurant-service/nats_client"
	"cap-club/restaurant-service/routes"
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
		go nats_client.UpdateRestaurants()
		go nats_client.DeleteRestaurant()
		go nats_client.UpdateRestaurant()
		router := gin.Default()
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
