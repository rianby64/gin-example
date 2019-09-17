package main_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var host string
var port string
var Error404NotFound = errors.New("404 Not found")
var mockRouter *gin.Engine

func requestResponse(method string, from string, body io.Reader) (interface{}, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(method, fmt.Sprintf("http://%s:%s/%s", host, port, from), nil)
	mockRouter.ServeHTTP(resp, req)
	if err != nil {
		return nil, err
	}

	if resp.Code == 404 {
		return nil, Error404NotFound
	}

	var response interface{}
	err = json.Unmarshal([]byte(resp.Body.String()), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
