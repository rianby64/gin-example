package client

// MockClient for testing
type MockClient struct {
}

// CreateEntry for testing
func (mc *MockClient) CreateEntry() (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ok",
	}, nil
}

// GetEntryList for testing
func (mc *MockClient) GetEntryList() (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ok",
	}, nil
}
