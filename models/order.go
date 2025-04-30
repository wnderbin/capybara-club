package models

type Order struct {
	Id              string `gorm:"primary_key" json:"id"`
	User_id         string `gorm:"not null" json:"user_id"`
	Restaurant_id   string `gorm:"not null" json:"restaurant_id"`
	Price           int    `gorm:"not null" json:"price"`
	Delivery_status string `gorm:"not null" json:"status"`
}
