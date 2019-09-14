package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupRouter the server
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// gin.H is a shortcut for map[string]interface{}
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"links": map[string]interface{}{
				"self": "/",
			},
		})
	})

	r.GET("/api/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"links": map[string]interface{}{
				"self": "api/articles",
			},
			"data": []interface{}{},
		})
	})

	return r
}

func main() {
	r := SetupRouter()

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
