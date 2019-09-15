package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	client1c "github.com/rianby64/gin-example/lib/1c-client"
)

// SetupClient1C constructs the client for main
func SetupClient1C() client1c.EntryClient {
	return client1c.NewEntryClient()
}

// SetupRouter the server
func SetupRouter(entryClient client1c.EntryClient) *gin.Engine {
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
		entries, _ := entryClient.GetEntryList()
		result := map[string]interface{}{
			"links": map[string]interface{}{
				"self": "api/articles",
			},
			"data": entries,
		}
		c.JSON(http.StatusOK, result)
	})

	r.POST("/api/articles", func(c *gin.Context) {
		result, _ := entryClient.CreateEntry()
		c.JSON(http.StatusOK, result)
	})

	return r
}

func main() {
	entryClient := SetupClient1C()
	r := SetupRouter(entryClient)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
