package clientserver

import "net/http"

// Interface interface
type Interface interface {
	Fetch(getData func() ([]interface{}, error)) (int, map[string]interface{})
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

// Fetch for Request
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

// NewRequest constructor
func NewRequest(self string) Interface {
	return &Request{
		self: self,
	}
}