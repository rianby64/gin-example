package main

import (
	"github.com/gin-gonic/gin"

	client1C "github.com/rianby64/gin-example/lib/1c-client"
	"github.com/rianby64/gin-example/lib/serializer"
)

// SetupClient1C constructs the client for main
func SetupClient1C() client1C.Interface {
	return client1C.New()
}

// SetupRouter the server
func SetupRouter(client1c client1C.Interface) *gin.Engine {
	base := gin.Default()
	api := serializer.NewEngine(base)

	api.HandleJSON("GET", "api/articles",
		func(c *serializer.Context) (statusCode int, jsonResponse interface{}) {
			return c.Fetch(func() ([]interface{}, error) {
				return client1c.GetEntryList()
			})
		},
	)

	api.HandleJSON("POST", "api/articles",
		func(c *serializer.Context) (statusCode int, jsonResponse interface{}) {
			return c.Create(func() (map[string]interface{}, error) {
				return client1c.CreateEntry()
			})
		},
	)

	return base
}

func main() {
	client1c := SetupClient1C()
	r := SetupRouter(client1c)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
