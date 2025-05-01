package nats_client

import (
	"cap-club/internal/database"
	"cap-club/internal/models"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func UpdateRestaurants() {
	nc, err := New(nats.DefaultURL)
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
