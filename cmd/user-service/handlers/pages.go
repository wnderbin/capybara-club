package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.html", nil)
}

func RegisterUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", nil)
}

func LoginUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
