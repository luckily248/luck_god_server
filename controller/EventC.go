package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func EvenCIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "/event/index.html", gin.H{
		"title": "Main website",
	})
}
