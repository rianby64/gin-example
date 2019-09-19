package encoder

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type source func() (interface{}, error)

// Interface interface
type Interface interface {
	Fetch(fn source) (int, interface{})
	Create(fn source) (int, interface{})
}

// Context struct
type Context struct {
	*gin.Context
}

// Fetch data
func (jar *Context) Fetch(fn source) (int, interface{}) {
	data, _ := fn()
	result := map[string]interface{}{
		"id":     "ID",
		"result": data,
	}
	return http.StatusOK, result
}

// Create resource
func (jar *Context) Create(fn source) (int, interface{}) {
	data, _ := fn()
	result := map[string]interface{}{
		"id":     "ID",
		"result": data,
	}
	return http.StatusCreated, result
}

// NewContext constructor
func NewContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}
