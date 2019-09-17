package serializer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Interface interface
type Interface interface {
	Fetch(getData func() ([]interface{}, error)) (int, map[string]interface{})
	Create(getData func() (map[string]interface{}, error)) (int, map[string]interface{})
}

// Context struct
type Context struct {
	*gin.Context
}

// Fetch data
func (jar *Context) Fetch(getData func() ([]interface{}, error)) (int, map[string]interface{}) {
	data, _ := getData()
	result := map[string]interface{}{
		"data": data,
	}
	return http.StatusOK, result
}

// Create resource
func (jar *Context) Create(getData func() (map[string]interface{}, error)) (int, map[string]interface{}) {
	data, _ := getData()
	result := map[string]interface{}{
		"data": map[string]interface{}{
			"type":       "photos",
			"id":         "550e8400-e29b-41d4-a716-446655440000",
			"attributes": data,
		},
	}
	return http.StatusCreated, result
}

// NewContext constructor
func NewContext(c *gin.Context) *Context {
	return &Context{
		c,
	}
}
