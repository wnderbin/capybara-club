package routes

import (
	"cap-club/cmd/order-service/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	r.GET("/orders", handlers.GetOrders)
}
