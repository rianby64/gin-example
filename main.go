package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rianby64/gin-example/encoder"
	client1C "github.com/rianby64/gin-example/lib/1c-client"
)

// SetupClient1C constructs the client for main
func SetupClient1C() client1C.Interface {
	return client1C.New()
}

// SetupRouter the server
func SetupRouter() (*gin.Engine, *encoder.Engine) {
	base := gin.Default()
	api := encoder.NewEngine(base)
	return base, api
}

func main() {
	r, e := SetupRouter()
	c := SetupClient1C()
	client1C.PlugToServer(c, e)

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
