package handlers

import (
	"cap-club/admin-service/config"
	"cap-club/admin-service/models"
	"cap-club/admin-service/utils"
	"cap-club/database"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddAdmin(c *gin.Context) {
	adm_name := c.Query("name")
	adm_email := c.Query("email")
	adm_pass := c.Query("password")

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

	hashed_password, err := utils.HashPassword(adm_pass)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not hash password..."})
		return
	}

	database.DB.Create(&models.Admin{
		Id:       uuid.NewString(),
		Name:     adm_name,
		Email:    adm_email,
		Password: hashed_password,
	})
	c.JSON(http.StatusCreated, gin.H{"message": "admin added successfully"})
}
