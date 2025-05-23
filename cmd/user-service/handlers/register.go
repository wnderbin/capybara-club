package handlers

import (
	"cap-club/internal/database"
	"cap-club/internal/models"
	"cap-club/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Register(c *gin.Context) {
	name := c.Request.FormValue("name")
	username := c.Request.FormValue("username")
	email := c.Request.FormValue("email")
	password := c.Request.FormValue("password")

	if name == "" || username == "" || email == "" || password == "" {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "you cannot register because you have specified empty fields."})
		return
	}

	hashed_password, err := utils.HashPassword(password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password..."})
		return
	}
	database.DB.Create(&models.User{
		Id:       uuid.NewString(),
		Name:     name,
		Username: username,

		Email:    email,
		Password: hashed_password,

		Created_at: time.Now(),
		Updated_at: time.Now(),
	})
	c.JSON(http.StatusCreated, gin.H{"message": "user registered successfully"})
}
