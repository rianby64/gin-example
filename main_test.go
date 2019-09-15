package main_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"

	main "github.com/rianby64/gin-example"
	client1c "github.com/rianby64/gin-example/lib/client1c"
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
	t.Run("1 - read empty", func(t *testing.T) {
		resp, err := requestResponse("GET", "api/articles", nil)
		if err != nil {
			t.Fatal(err)
		}

		expected, _ := mockClient1C.GetEntryList()
		assert.Equal(t, expected, resp)
	})

	t.Run("2 - post -> OK", func(t *testing.T) {
		resp, err := requestResponse("POST", "api/articles", nil)
		if err != nil {
			t.Fatal(err)
		}

		expected, _ := mockClient1C.CreateEntry()
		assert.Equal(t, expected, resp)
	})
}
