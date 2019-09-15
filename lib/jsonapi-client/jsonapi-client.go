package client

import "net/http"

// JSONAPI interface
type JSONAPI interface {
	Fetch(entries []interface{}) (int, map[string]interface{})
}

// JSONAPIRequest struct
type JSONAPIRequest struct {
}

// JSONAPIResponse struct
type JSONAPIResponse struct {
	data   interface{}
	errors []interface{}
	meta   map[string]interface{}

	jsonapi  string
	links    map[string]interface{}
	included []interface{}
}

// Fetch for JSONAPIRequest
func (*JSONAPIRequest) Fetch(entries []interface{}) (int, map[string]interface{}) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": entries,
	}
	return http.StatusOK, result
}

// NewJSONAPIRequest constructor
func NewJSONAPIRequest() JSONAPI {
	return &JSONAPIRequest{}
}
