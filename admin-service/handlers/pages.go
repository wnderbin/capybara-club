package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUserForm(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
