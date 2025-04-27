package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/nats_client"
	"cap-club/restaurant-service/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func UpdateRestaurant(c *gin.Context) {
	cookie, err := c.Cookie("jwt-admin")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no token found"})
		return
	}

	conf := config.MustLoad()

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(cookie, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JWTKey), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	id := c.Query("id")
	name := c.Query("name")
	address := c.Query("address")
	email := c.Query("email")
	phone := c.Query("phone")
	description := c.Query("description")

	nc, err := nats_client.New(nats.DefaultURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "nuts error"})
		return
	}
	defer nc.Close()
	msg := models.Restaurant{
		Id:          id,
		Name:        name,
		Address:     address,
		Email:       email,
		PhoneNumber: phone,
		Created_at:  time.Now(),
		Description: description,
	}
	jsonMsg, err := json.Marshal(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "json marshaling failed"})
		return
	}
	err = nc.Conn.Publish("update.restaurant", jsonMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error publishing"})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "updated"})
	}
}
