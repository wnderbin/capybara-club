package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MainPageHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "main_page.html", nil)
}
