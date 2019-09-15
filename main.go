package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	client1C "github.com/rianby64/gin-example/lib/1c-client"
	clientJSONAPI "github.com/rianby64/gin-example/lib/jsonapi-client"
)

// SetupClient1C constructs the client for main
func SetupClient1C() client1C.EntryClient {
	return client1C.NewEntryClient()
}

// DummyMiddleware first md
func DummyMiddleware(c *gin.Context) {
	// this fn is being executed before any handler

	// Pass on to the next-in-chain
	c.Next()
}

// SetupRouter the server
func SetupRouter(entryClient client1C.EntryClient) *gin.Engine {
	api := gin.Default()
	api.Use(DummyMiddleware)

	// gin.H is a shortcut for map[string]interface{}
	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"links": map[string]interface{}{
				"self": "/",
			},
		})
	})

	api.GET("/api/articles", func(c *gin.Context) {
		jsonapi := clientJSONAPI.NewJSONAPIRequest()
		entries, _ := entryClient.GetEntryList()
		status, result := jsonapi.Fetch(entries)
		c.JSON(status, result)
	})

	api.POST("/api/articles", func(c *gin.Context) {
		result, _ := entryClient.CreateEntry()
		c.JSON(http.StatusOK, result)
	})

	return api
}

func main() {
	entryClient := SetupClient1C()
	r := SetupRouter(entryClient)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
