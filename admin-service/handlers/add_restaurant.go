package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/nats_client"
	"cap-club/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
)

func AddRestaurant(c *gin.Context) {
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

	name := c.Request.FormValue("name")
	address := c.Request.FormValue("address")
	email := c.Request.FormValue("email")
	phone := c.Request.FormValue("phone")
	description := c.Request.FormValue("description")

	nc, err := nats_client.New(nats.DefaultURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "nuts error"})
		return
	}
	defer nc.Close()
	msg := models.Restaurant{
		Id:          uuid.NewString(),
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
	err = nc.Conn.Publish("add.restaurant", jsonMsg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error publishing"})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "published"})
	}
}
