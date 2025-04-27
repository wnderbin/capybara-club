package models

type Restaurant struct {
	Id          string `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not null;unique" json:"name"`
	Address     string `gorm:"not null;unique" json:"address"`
	Email       string `gorm:"not null" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;not null" json:"phone"`
	Created_at  string `gorm:"not null" json:"created_at"`
	Description string `gorm:"not null" json:"description"`
}
