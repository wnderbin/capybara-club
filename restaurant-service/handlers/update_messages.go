package handlers

import (
	"cap-club/restaurant-service/database"
	"cap-club/restaurant-service/models"
	"cap-club/restaurant-service/nats_client"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func UpdateRestaurants() {
	nc, err := nats_client.New(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	defer nc.Close()
	_, err = nc.Conn.Subscribe("add.restaurant", func(m *nats.Msg) {
		var recievedRestaurant models.Restaurant
		err = json.Unmarshal(m.Data, &recievedRestaurant)
		if err != nil {
			database.Log.Error("error unmarshaling from json")
			return
		}
		database.Log.Info("Recieved restaurant")
		err = database.DB.Create(&recievedRestaurant).Error
		if err != nil {
			database.Log.Error("Error inserting into database")
			return
		}
	})
	if err != nil {
		database.Log.Error("subscribe error")
		return
	}
	select {}
}
