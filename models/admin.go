package models

type Admin struct {
	Id       string `gorm:"primary_key" json:"id"`
	Name     string `gorm:"not null;unique" json:"name"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}
