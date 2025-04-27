package handlers

import (
	"cap-club/restaurant-service/database"
	"cap-club/restaurant-service/models"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func update_restaurants() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	defer nc.Close()
	_, err = nc.Subscribe("add-restaurant", func(m *nats.Msg) {
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
