package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/database"
	"cap-club/admin-service/models"
	"cap-club/admin-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var admin models.Admin

	admin_name := c.Request.FormValue("name")
	admin_password := c.Request.FormValue("password")
	fmt.Println(admin_name, admin_password)

	err := database.DB.Where("name = ?", admin_name).Take(&admin).Error
	fmt.Println(admin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "admin with this name does not exist"})
		return
	}

	if passwordStatus := utils.CheckHashedPassword(admin_password, admin.Password); !passwordStatus {
		c.JSON(http.StatusBadRequest, gin.H{"error": "incorrect password specified"})
		return
	}

	conf := config.MustLoad()

	token, err := utils.GenerateJWT(admin.Name, conf)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "jwt-admin",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		MaxAge:   300,
	})

	c.JSON(http.StatusAccepted, gin.H{"message": "your token was saved in cookies and its lifetime is about 5 minutes, after this time you will need to log in to your admin account again."})
}
