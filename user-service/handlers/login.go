package handlers

import (
	"cap-club/user-service/config"
	"cap-club/user-service/database"
	"cap-club/user-service/models"
	"cap-club/user-service/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user models.User

	user_username := c.Request.FormValue("username")
	user_password := c.Request.FormValue("password")

	err := database.DB.Where("username = ?", user_username).First(&user).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user with this username does not exist"})
		return
	}

	if passwordStatus := utils.CheckHashedPassword(user_password, user.Password); !passwordStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect password specified"})
		return
	}

	conf := config.MustLoad()

	token, err := utils.GenerateJWT(user.Username, conf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}
