package clientserver

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Engine the JSONAPI extension for gin engine
type Engine struct {
	*gin.Engine
}

// Handler signature
type Handler func(jsonapi Interface, ctx *gin.Context) (statusCode int, jsonResponse interface{})

// Middleware formats the responses by JSONAPI
func Middleware(ctx *gin.Context) {
	ctx.Next()
	jsonapi := NewRequest(ctx.Request.URL.String())
	handlerPtr, exists := ctx.Get("Handler")
	if exists {
		handler, ok := handlerPtr.(Handler)
		if !ok {
			log.Println("cannot convert a handler into Handler")
		} else {
			statusCode, jsonResponse := handler(jsonapi, ctx)
			ctx.JSON(statusCode, jsonResponse)
		}
	}
}

// HandleJSONAPI registers a method and URL to match a handler
func (e *Engine) HandleJSONAPI(method, URL string, handler Handler) {
	e.Handle(method, URL, func(ctx *gin.Context) {
		ctx.Set("Handler", handler)
	})
}

// NewEngine constructor
func NewEngine(api *gin.Engine) *Engine {
	e := &Engine{api}
	e.Use(Middleware)
	return e
}
