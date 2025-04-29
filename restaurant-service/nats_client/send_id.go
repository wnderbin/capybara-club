package nats_client

import (
	"cap-club/database"
	"cap-club/restaurant-service/models"

	"github.com/nats-io/nats.go"
)

func SendId() {
	nc, err := New(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	var restaurantID string
	defer nc.Close()
	_, err = nc.Conn.Subscribe("get.restaurant.id", func(m *nats.Msg) {
		var recievedName string
		var restaurant models.Restaurant
		recievedName = string(m.Data)
		database.Log.Info("Received name")
		err = database.DB.Where("name = ?", recievedName).Find(&restaurant).Error
		if err != nil {
			database.Log.Error("Error deleting from database")
			return
		}
		restaurantID = restaurant.Id
		err = nc.Conn.Publish("send.restaurant.id", []byte(restaurantID))
		if err != nil {
			database.Log.Error("failed to send id")
			return
		}
	})
	if err != nil {
		database.Log.Error("subscribe error")
		return
	}
	select {}
}
