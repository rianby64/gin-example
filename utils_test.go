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
var ErrorNot200OK = errors.New("Status is not 200")
var mockRouter *gin.Engine

func requestResponse(method string, from string, body io.Reader) (interface{}, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest(method, fmt.Sprintf("http://%s:%s/%s", host, port, from), nil)
	mockRouter.ServeHTTP(resp, req)
	if err != nil {
		return nil, err
	}

	if resp.Code != 200 {
		return nil, ErrorNot200OK
	}

	var response interface{}
	err = json.Unmarshal([]byte(resp.Body.String()), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
