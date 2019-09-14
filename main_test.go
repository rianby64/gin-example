package main_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	main "github.com/rianby64/gin-example"
)

var host string
var port string
var ErrorNot200OK = errors.New("Status is not 200")
var mockRouter *gin.Engine

func init() {
	err := godotenv.Load()
	if err == nil {
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
	}
	fmt.Printf("Running host=%s\n", host)
	mockRouter = main.SetupRouter()
}

func requestResponse(from string) (interface{}, error) {
	resp := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s/%s", host, port, from), nil)
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

func Test_check_connection(t *testing.T) {
	_, err := requestResponse("")
	if err != nil {
		t.Fatal(err)
	}
}

func Test_articles_CRUD__OK(t *testing.T) {
	t.Run("1 - read empty", func(t *testing.T) {
		resp, err := requestResponse("api/articles")
		if err != nil {
			t.Fatal(err)
		}

		expected := map[string]interface{}{
			"links": map[string]interface{}{
				"self": "api/articles",
			},
			"data": []interface{}{},
		}
		assert.Equal(t, expected, resp)
	})
}
