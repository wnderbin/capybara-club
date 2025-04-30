package handlers

import (
	"cap-club/database"
	"cap-club/models"
	"cap-club/user-service/config"
	"cap-club/utils"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password specified"})
		return
	}

	conf := config.MustLoad()

	token, err := utils.GenerateJWTUser(user.Username, conf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{ // we set cookies to have access to the jwt token for further manipulations
		Name:     "jwt-token", // cookie name, this name will be used to store and send the cookie to the server
		Value:    token,       // cookie value
		Path:     "/",         // the path for which the cookie will be available, in this case the cookie will be available for all routes
		HttpOnly: true,        // prevents cookies from being accessed via javascript, helps protect cookies from XSS attacks
		Secure:   false,       // if set to true, it means that cookies will only be transmitted over secure connections
		MaxAge:   300,         // cookie lifetime, in seconds
	})

	c.JSON(http.StatusAccepted, gin.H{"message": "your token was saved in cookies and its lifetime is about 5 minutes, after this time you will need to log in to your account again."})
}
