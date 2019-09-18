package plug

import (
	"github.com/rianby64/gin-example/encoder"
	client1C "github.com/rianby64/gin-example/lib/1c-client"
)

// define here all these handlers...

func getArticles(client1c client1C.Interface) encoder.Handler {
	return func(c *encoder.Context) (statusCode int, jsonResponse interface{}) {
		return c.Fetch(func() ([]interface{}, error) {
			return client1c.GetEntryList()
		})
	}
}

func postArticles(client1c client1C.Interface) encoder.Handler {
	return func(c *encoder.Context) (statusCode int, jsonResponse interface{}) {
		return c.Create(func() (map[string]interface{}, error) {
			return client1c.CreateEntry()
		})
	}
}

// Client1C links this client with server
func Client1C(client1c client1C.Interface, e *encoder.Engine) {
	g := e.GroupJSON("api")
	g.HandleJSON("GET", "articles", getArticles(client1c))
	g.HandleJSON("POST", "articles", postArticles(client1c))
}
