package handlers

import (
	"cap-club/user-service/config"
	"cap-club/user-service/database"
	"cap-club/user-service/models"
	"cap-club/user-service/utils"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateUser(c *gin.Context) {
	var user models.User

	name := c.Query("name")
	username := c.Query("username")
	email := c.Query("email")
	password := c.Query("password")

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

	err = database.DB.Where("username = ?", userUsername).First(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find user"})
		return
	}

	if name != "" || username != "" || email != "" || password != "" {
		user.Name = name
		user.Username = username
		user.Email = email
		hashed_password, err := utils.HashPassword(password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password..."})
			return
		}
		user.Password = hashed_password
		user.Updated_at = time.Now()
	} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "the changes cannot be applied because you have specified empty fields"})
		return
	}

	err = database.DB.Save(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update user"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "updated successfully, since you have changed your details, you need to register again"})
}
