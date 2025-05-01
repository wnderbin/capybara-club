package main

import (
	"cap-club/cmd/order-service/config"
	"cap-club/cmd/order-service/nats_client"
	"cap-club/internal/database"
	"cap-club/internal/migrator"
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
