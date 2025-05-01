package routes

import (
	"cap-club/cmd/restaurant-service/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// --- GET ---
	r.GET("/restaurants", handlers.GetRestautants)
}
