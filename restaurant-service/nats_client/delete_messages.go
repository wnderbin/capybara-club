package nats_client

import (
	"cap-club/database"
	"cap-club/restaurant-service/models"

	"github.com/nats-io/nats.go"
)

func DeleteRestaurant() {
	nc, err := New(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	defer nc.Close()
	_, err = nc.Conn.Subscribe("delete.restaurant", func(m *nats.Msg) {
		var recievedName string
		var restaurant models.Restaurant
		recievedName = string(m.Data)
		database.Log.Info("Recieved restaurant")
		err = database.DB.Where("name = ?", recievedName).Delete(&restaurant).Error
		if err != nil {
			database.Log.Error("Error deleting from database")
			return
		}
	})
	if err != nil {
		database.Log.Error("subscribe error")
		return
	}
	select {}
}
