package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	main "github.com/rianby64/gin-example"
	client1c "github.com/rianby64/gin-example/lib/1c-client"
)

var mockClient1C *client1c.MockClient

func init() {
	err := godotenv.Load()
	if err == nil {
		host = os.Getenv("HOST")
		port = os.Getenv("PORT")
	}
	fmt.Printf("Running host=%s\n", host)
	mockClient1C = &client1c.MockClient{}
	mockRouter = main.SetupRouter(mockClient1C)
}

func Test_check_connection(t *testing.T) {
	_, err := requestResponse("GET", "", nil)
	if err != nil {
		t.Fatal(err)
	}
}

func Test_articles_CRUD__OK(t *testing.T) {
	t.Run("get list", func(t *testing.T) {
		self := "api/articles"
		resp, err := requestResponse("GET", self, nil)
		if err != nil {
			t.Fatal(err)
		}

		data, _ := mockClient1C.GetEntryList()
		expected := map[string]interface{}{
			"links": map[string]interface{}{
				"self": fmt.Sprintf("http://%s:%s/%s", host, port, self),
			},
			"data": data,
		}
		assert.Equal(t, expected, resp)
	})

	/*
		t.Run("get", func(t *testing.T) {
			resp, err := requestResponse("GET", "api/articles/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			expected, _ := mockClient1C.GetEntryList()
			assert.Equal(t, expected, resp)
		})

		t.Run("post", func(t *testing.T) {
			resp, err := requestResponse("POST", "api/articles", nil)
			if err != nil {
				t.Fatal(err)
			}

			expected, _ := mockClient1C.CreateEntry()
			assert.Equal(t, expected, resp)
		})

		t.Run("patch", func(t *testing.T) {
			resp, err := requestResponse("PATCH", "api/articles/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			expected, _ := mockClient1C.CreateEntry()
			assert.Equal(t, expected, resp)
		})

		t.Run("delete", func(t *testing.T) {
			resp, err := requestResponse("DELETE", "api/articles/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			expected, _ := mockClient1C.CreateEntry()
			assert.Equal(t, expected, resp)
		})
	*/
}
