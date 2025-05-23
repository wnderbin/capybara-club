package handlers

import (
	"cap-club/cmd/user-service/config"
	"cap-club/internal/database"
	"cap-club/internal/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
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

	err = database.DB.Where("username = ?", userUsername).Delete(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete user"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "the user was successfully deleted"})
}
