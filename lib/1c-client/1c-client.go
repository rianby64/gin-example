package client

// Interface manages the proxy to bring Entries from somewhere
type Interface interface {
	GetEntryList() (interface{}, error)
	GetEntry() (interface{}, error)
	CreateEntry() (interface{}, error)
	UpdateEntry() (interface{}, error)
	DeleteEntry() (interface{}, error)
}

type client struct {
}

// GetEntryList retreives from 1C the info
func (c *client) GetEntryList() (interface{}, error) {
	return []interface{}{}, nil
}

// GetEntry retreives from 1C the info
func (c *client) GetEntry() (interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// CreateEntry creates an entry for 1C
func (c *client) CreateEntry() (interface{}, error) {
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
func (c *client) UpdateEntry() (interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// DeleteEntry retreives from 1C the info
func (c *client) DeleteEntry() (interface{}, error) {
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
