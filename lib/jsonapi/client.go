package clientserver

import "net/http"

// Interface interface
type Interface interface {
	Fetch(getData func() ([]interface{}, error)) (int, map[string]interface{})
	Create(getData func() (map[string]interface{}, error)) (int, map[string]interface{})
}

// Request struct
type Request struct {
	self string
}

/*
// Response struct
type Response struct {
	data   interface{}
	errors []interface{}
	meta   map[string]interface{}

	Interface  string
	links    map[string]interface{}
	included []interface{}
}
*/

// Fetch data
func (jar *Request) Fetch(getData func() ([]interface{}, error)) (int, map[string]interface{}) {
	data, _ := getData()
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": jar.self,
		},
		"data": data,
	}
	return http.StatusOK, result
}

// Create resource
func (jar *Request) Create(getData func() (map[string]interface{}, error)) (int, map[string]interface{}) {
	data, _ := getData()
	result := map[string]interface{}{
		"data": map[string]interface{}{
			"type":       "photos",
			"id":         "550e8400-e29b-41d4-a716-446655440000",
			"attributes": data,
			"links": map[string]interface{}{
				"self": jar.self,
			},
		},
	}
	return http.StatusCreated, result
}

// NewRequest constructor
func NewRequest(self string) Interface {
	return &Request{
		self: self,
	}
}
