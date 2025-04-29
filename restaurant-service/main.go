package main

import (
	"cap-club/database"
	"cap-club/migrator"
	"cap-club/restaurant-service/config"
	"cap-club/restaurant-service/nats_client"
	"cap-club/restaurant-service/routes"
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
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
		go nats_client.UpdateRestaurants()
		go nats_client.DeleteRestaurant()
		go nats_client.UpdateRestaurant()
		go nats_client.SendId()
		router := gin.Default()
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
