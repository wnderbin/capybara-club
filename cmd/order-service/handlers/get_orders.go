package handlers

import (
	"cap-club/internal/database"
	"cap-club/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrders(c *gin.Context) {
	user_id := c.Query("id")
	var orders []models.Order
	err := database.DB.Where("user_id = ?", user_id).Find(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "find restaurants error"})
		return
	}
	c.JSON(http.StatusOK, orders)
}
