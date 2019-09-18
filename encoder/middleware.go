package encoder

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Engine is the extension-encoder for *gin.Engine
type Engine struct {
	*gin.Engine
}

// RouterGroup is the extension-encoder for gin RouterGroup
type RouterGroup struct {
	*gin.RouterGroup
}

// Handler signature
type Handler func(ce *Context) (statusCode int, jsonResponse interface{})

// Middleware formats the responses by the encoder
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
func (rg *RouterGroup) HandleJSON(method, URL string, handler Handler) {
	rg.Handle(method, URL, func(c *gin.Context) {
		c.Set("Handler", handler)
	})
}

// HandleJSON registers a method and URL to match a handler
func (e *Engine) HandleJSON(method, URL string, handler Handler) {
	e.Handle(method, URL, func(c *gin.Context) {
		c.Set("Handler", handler)
	})
}

// GroupJSON what a heck
func (e *Engine) GroupJSON(URL string) *RouterGroup {
	rg := RouterGroup{e.Group(URL)}
	return &rg
}

// NewEngine constructor
func NewEngine(api *gin.Engine) *Engine {
	e := &Engine{api}
	e.Use(Middleware)
	return e
}
