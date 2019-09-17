package serializer

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Engine the Serializer extension for gin engine
type Engine struct {
	*gin.Engine
}

// Handler signature
type Handler func(ce *Context) (statusCode int, jsonResponse interface{})

// Middleware formats the responses by Serializer
func Middleware(c *gin.Context) {
	c.Next()
	handlerPtr, exists := c.Get("Handler")
	if exists {
		handler, ok := handlerPtr.(Handler)
		if !ok {
			log.Println("cannot convert a handler into Handler")
		} else {
			ce := NewContext(c)
			statusCode, jsonResponse := handler(ce)
			c.JSON(statusCode, jsonResponse)
		}
	}
}

// HandleJSON registers a method and URL to match a handler
func (e *Engine) HandleJSON(method, URL string, handler Handler) {
	e.Handle(method, URL, func(c *gin.Context) {
		c.Set("Handler", handler)
	})
}

// NewEngine constructor
func NewEngine(api *gin.Engine) *Engine {
	e := &Engine{api}
	e.Use(Middleware)
	return e
}
