package models

import "time"

type User struct {
	Id       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null" json:"name"`
	Username string `gorm:"not null;unique" json:"username"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`

	Created_at time.Time `gorm:"not null" json:"created_at"`
	Updated_at time.Time `gorm:"not null" json:"updated_at"`
}
