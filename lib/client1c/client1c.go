package client

// EntryClient manages the proxy to bring Entries from somewhere
type EntryClient interface {
	GetEntryList() (map[string]interface{}, error)
	GetEntry() (map[string]interface{}, error)
	CreateEntry() (map[string]interface{}, error)
	UpdateEntry() (map[string]interface{}, error)
	DeleteEntry() (map[string]interface{}, error)
}

type client struct {
}

// GetEntryList retreives from 1C the info
func (ec *client) GetEntryList() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// GetEntry retreives from 1C the info
func (ec *client) GetEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// CreateEntry creates an entry for 1C
func (ec *client) CreateEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"data": map[string]interface{}{
			"type": "photos",
			"id":   "550e8400-e29b-41d4-a716-446655440000",
			"attributes": map[string]interface{}{
				"title": "Ember Hamster",
				"src":   "http://example.com/images/productivity.png",
			},
			"links": map[string]interface{}{
				"self": "http://example.com/photos/550e8400-e29b-41d4-a716-446655440000",
			},
		},
	}
	return result, nil
}

// UpdateEntry retreives from 1C the info
func (ec *client) UpdateEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// DeleteEntry retreives from 1C the info
func (ec *client) DeleteEntry() (map[string]interface{}, error) {
	result := map[string]interface{}{
		"links": map[string]interface{}{
			"self": "api/articles",
		},
		"data": []interface{}{},
	}
	return result, nil
}

// NewEntryClient is the constructor for this proxy
func NewEntryClient() EntryClient {
	c := client{}
	return &c
}
