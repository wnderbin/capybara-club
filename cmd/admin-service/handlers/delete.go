package handlers

import (
	"cap-club/cmd/admin-service/config"
	"cap-club/internal/database"
	"cap-club/internal/models"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DeleteAdmin(c *gin.Context) {
	var admin models.Admin

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
	}

	admName := claims["admin"].(string)
	if admName == conf.AdminUsername {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "can't delete root administrator"})
		return
	}
	err = database.DB.Where("name = ?", admName).Delete(&admin).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to delete admin"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "the admin was successfully deleted"})
}
