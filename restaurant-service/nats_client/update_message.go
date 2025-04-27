package nats_client

import (
	"cap-club/restaurant-service/database"
	"cap-club/restaurant-service/models"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

func UpdateRestaurant() {
	nc, err := New(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	defer nc.Close()
	_, err = nc.Conn.Subscribe("update.restaurant", func(m *nats.Msg) {
		var recievedRestaurant models.Restaurant
		var restaurant models.Restaurant
		err = json.Unmarshal(m.Data, &recievedRestaurant)
		if err != nil {
			database.Log.Error("error unmarshaling from json")
			return
		}
		database.Log.Info("Recieved restaurant")
		err = database.DB.Where("id = ?", recievedRestaurant.Id).Find(&restaurant).Error
		if err != nil {
			database.Log.Error("Error while getting restaurant")
			return
		}
		restaurant.Name = recievedRestaurant.Name
		restaurant.Address = recievedRestaurant.Address
		restaurant.Email = recievedRestaurant.Email
		restaurant.PhoneNumber = recievedRestaurant.PhoneNumber
		restaurant.Created_at = recievedRestaurant.Created_at
		restaurant.Description = recievedRestaurant.Description

		err = database.DB.Save(&restaurant).Error
		if err != nil {
			database.Log.Error("Error updating in database")
			return
		}
	})
	if err != nil {
		database.Log.Error("subscribe error")
		return
	}
	select {}
}
