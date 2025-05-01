package routes

import (
	"cap-club/cmd/admin-service/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// --- GET ---
	r.GET("/login/", handlers.LoginUserForm)
	r.GET("/add_restaurant/", handlers.AddRestaurantForm)
	// --- POST ---
	r.POST("/login/postform", handlers.Login)
	r.POST("/admin", handlers.AddAdmin)
	r.POST("/add_restaurant/postform", handlers.AddRestaurant)
	// --- PUT ---
	r.PUT("/admin", handlers.UpdateAdmin)
	r.PUT("/restaurant", handlers.UpdateRestaurant)
	// --- DELETE ---
	r.DELETE("/admin", handlers.DeleteAdmin)
	r.DELETE("/restaurant", handlers.DeleteRestaurant)
}
