package routes

import (
	"cap-club/admin-service/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// --- GET ---
	r.GET("/login/", handlers.LoginUserForm)
	// --- POST ---
	r.POST("/login/postform", handlers.Login)
	r.POST("/admin", handlers.AddAdmin)
	// --- PUT ---
	r.PUT("/admin", handlers.UpdateAdmin)
	// --- DELETE ---
	r.DELETE("/admin", handlers.DeleteAdmin)
}
