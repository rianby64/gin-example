package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	client1c "github.com/rianby64/gin-example/lib/1c-client"
)

func Test_client1c_CRUD__OK(t *testing.T) {
	client := client1c.New()
	t.Run("get entry list", func(t *testing.T) {
		resp, err := client.GetEntryList()
		if err != nil {
			t.Log(err)
			return
		}

		arr := resp.([]interface{})
		expected := 8
		assert.Equal(t, expected, len(arr))
	})

	t.Run("create entry", func(t *testing.T) {
		resp, err := client.CreateEntry()
		if err != nil {
			t.Fatal(err)
		}

		expected := map[string]interface{}{
			"type": "photos",
			"id":   "550e8400-e29b-41d4-a716-446655440000",
			"attributes": map[string]interface{}{
				"title": "Ember Hamster",
				"src":   "http://example.com/images/productivity.png",
			},
		}
		assert.Equal(t, expected, resp)
	})
}
