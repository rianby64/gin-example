package client

import "github.com/rianby64/gin-example/encoder"

// Interface manages the proxy to bring Entries from somewhere
type Interface interface {
	GetEntryList() ([]interface{}, error)
	GetEntry() (map[string]interface{}, error)
	CreateEntry() (map[string]interface{}, error)
	UpdateEntry() (map[string]interface{}, error)
	DeleteEntry() (map[string]interface{}, error)
}

type client struct {
}

// GetEntryList retreives from 1C the info
func (c *client) GetEntryList() ([]interface{}, error) {
	return []interface{}{}, nil
}

// GetEntry retreives from 1C the info
func (c *client) GetEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// CreateEntry creates an entry for 1C
func (c *client) CreateEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"type": "photos",
		"id":   "550e8400-e29b-41d4-a716-446655440000",
		"attributes": map[string]interface{}{
			"title": "Ember Hamster",
			"src":   "http://example.com/images/productivity.png",
		},
	}
	return result, nil
}

// UpdateEntry retreives from 1C the info
func (c *client) UpdateEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// DeleteEntry retreives from 1C the info
func (c *client) DeleteEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// New is the constructor for this proxy
func New() Interface {
	client1c := client{}
	return &client1c
}

// PlugToServer links this client with server
func PlugToServer(client1c Interface, e *encoder.Engine) {

	g := e.GroupJSON("api")

	g.HandleJSON("GET", "articles",
		func(c *encoder.Context) (statusCode int, jsonResponse interface{}) {
			return c.Fetch(func() ([]interface{}, error) {
				return client1c.GetEntryList()
			})
		},
	)

	g.HandleJSON("POST", "articles",
		func(c *encoder.Context) (statusCode int, jsonResponse interface{}) {
			return c.Create(func() (map[string]interface{}, error) {
				return client1c.CreateEntry()
			})
		},
	)
}
