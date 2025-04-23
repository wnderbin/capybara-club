package routes

import (
	"cap-club/user-service/handlers"

	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	// --- GET ---
	r.GET("/main", handlers.MainPageHandler)
	r.GET("/register/", handlers.RegisterUserForm)
	r.GET("/login/", handlers.LoginUserForm)
	r.GET("/user", handlers.GetUser)

	// --- POST ---

	r.POST("/register/postform", handlers.Register)
	r.POST("/login/postform", handlers.Login)

	// --- DELETE ---

	r.DELETE("/user", handlers.DeleteUser)

	// --- PUT ---

	r.PUT("/user", handlers.UpdateUser) // user?name=_&username=_&email=_&password=_
}
