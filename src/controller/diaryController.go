package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Runner() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/", func(context *gin.Context) {

	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
