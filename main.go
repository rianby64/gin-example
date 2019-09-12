package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hey",
			"status":  http.StatusOK,
		})
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
