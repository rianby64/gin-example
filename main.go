package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	client1C "github.com/rianby64/gin-example/lib/1c-client"
	jsonAPI "github.com/rianby64/gin-example/lib/jsonapi"
)

// SetupClient1C constructs the client for main
func SetupClient1C() client1C.Interface {
	return client1C.New()
}

// SetupRouter the server
func SetupRouter(client1c client1C.Interface) *gin.Engine {
	api := gin.Default()
	api4JSON := jsonAPI.NewEngine(api)

	api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"links": map[string]interface{}{
				"self": "/",
			},
		})
	})

	api4JSON.HandleJSONAPI("GET", "api/articles",
		func(jsonapi jsonAPI.Interface, ctx *gin.Context) (statusCode int, jsonResponse interface{}) {
			return jsonapi.Fetch(func() ([]interface{}, error) {
				return client1c.GetEntryList()
			})
		},
	)

	api4JSON.HandleJSONAPI("POST", "api/articles",
		func(jsonapi jsonAPI.Interface, ctx *gin.Context) (statusCode int, jsonResponse interface{}) {
			return jsonapi.Create(func() (map[string]interface{}, error) {
				return client1c.CreateEntry()
			})
		},
	)

	return api
}

func main() {
	client1c := SetupClient1C()
	r := SetupRouter(client1c)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
