package handlers

import (
	"cap-club/database"
	"cap-club/models"
	"cap-club/user-service/config"
	"cap-club/user-service/nats_client"
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
)

func CreateOrder(c *gin.Context) {
	var user models.User
	cookie, err := c.Cookie("jwt-token")
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
	userUsername := claims["username"].(string)
	err = database.DB.Where("username = ?", userUsername).Find(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find user"})
		return
	}
	userID := user.Id
	restaurantName := c.Query("restaurant_name")
	nc, err := nats_client.New(nats.DefaultURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "nuts error"})
		return
	}
	defer nc.Close()
	msg1 := restaurantName
	err = nc.Conn.Publish("get.restaurant.id", []byte(msg1))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "failed to get id")
		return
	}

	// ---

	idChannel := make(chan string)
	_, err = nc.Conn.Subscribe("send.restaurant.id", func(m *nats.Msg) {
		message := string(m.Data)
		database.Log.Info("Received id: %s", string(message))
		idChannel <- message
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "subscription error"})
		return
	}

	select {
	case restaurantID := <-idChannel:
		database.Log.Info("Restaurant ID received: %s", restaurantID)
		msg2 := models.Message{
			UserId:       userID,
			RestaurantId: restaurantID,
		}
		jsonMsg, err := json.Marshal(msg2)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "json marshaling failed"})
			return
		}

		err = nc.Conn.Publish("create.order", jsonMsg)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "not published"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"status": "published"})
	case <-time.After(15 * time.Second):
		c.JSON(http.StatusGatewayTimeout, gin.H{"error": "timeout waiting for restaurant ID"})
		return
	}
}
