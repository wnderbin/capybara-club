package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/database"
	"cap-club/models"
	"cap-club/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UpdateAdmin(c *gin.Context) {
	var admin models.Admin

	new_name := c.Query("name")
	new_email := c.Query("email")
	new_password := c.Query("password")

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
	hashed_password, err := utils.HashPassword(new_password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password..."})
		return
	}
	admName := claims["admin"].(string)
	if admName == conf.AdminUsername {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "root administrator cannot be changed"})
		return
	}

	err = database.DB.Where("name = ?", admName).First(&admin).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to find admin"})
		return
	}

	admin.Name = new_name
	admin.Email = new_email
	admin.Password = hashed_password

	err = database.DB.Save(&admin).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to update admin"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "updated successfully, since you have changed your details, you need to login again"})
}
