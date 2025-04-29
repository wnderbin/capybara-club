package nats_client

import (
	"cap-club/database"
	"cap-club/order-service/models"
	"encoding/json"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func CreateOrder() {
	nc, err := New(nats.DefaultURL)
	if err != nil {
		database.Log.Error("[nuts] nuts error")
		return
	}
	defer nc.Close()
	_, err = nc.Conn.Subscribe("create.order", func(m *nats.Msg) {
		var receivedMessage models.Message
		err = json.Unmarshal(m.Data, &receivedMessage)
		if err != nil {
			database.Log.Error("error unmarshaling from json")
			return
		}
		database.Log.Info("Recieved message")
		err = database.DB.Create(&models.Order{
			Id:              uuid.NewString(),
			User_id:         receivedMessage.UserId,
			Restaurant_id:   receivedMessage.RestaurantId,
			Price:           100,
			Delivery_status: "active",
		}).Error
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
