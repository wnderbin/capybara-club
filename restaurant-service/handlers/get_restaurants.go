package handlers

import (
	"cap-club/database"
	"cap-club/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRestautants(c *gin.Context) {
	var restautants []models.Restaurant
	err := database.DB.Find(&restautants).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find restaurants error"})
		return
	}
	c.JSON(http.StatusOK, restautants)
}
