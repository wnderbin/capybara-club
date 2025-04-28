package main

import (
	"cap-club/database"
	"cap-club/migrator"
	"cap-club/user-service/config"
	"cap-club/user-service/routes"
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

		router := gin.Default()
		router.LoadHTMLGlob("user-service/ui/html/*")
		routes.Router(router)
		router.Run(fmt.Sprintf("%s:%d", conf.Address, conf.Port))
	}
}
