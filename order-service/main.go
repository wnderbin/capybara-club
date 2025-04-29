package main

import (
	"cap-club/database"
	"cap-club/migrator"
	"cap-club/order-service/config"
	"cap-club/order-service/nats_client"
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

		go nats_client.CreateOrder()

		router := gin.Default()
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
