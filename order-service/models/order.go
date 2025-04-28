package models

type Order struct {
	Id              string `json:"id"`
	User_id         string `json:"user_id"`
	Restaurant_id   string `json:"restaurant_id"`
	Price           int    `json:"price"`
	Delivery_status string `json:"status"`
}
