package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/nats_client"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func DeleteRestaurant(c *gin.Context) {
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

	name := c.Query("name")

	nc, err := nats_client.New(nats.DefaultURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "nuts error"})
		return
	}
	defer nc.Close()

	err = nc.Conn.Publish("delete.restaurant", []byte(name))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error deleting"})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{"status": "deleted"})
	}
}
