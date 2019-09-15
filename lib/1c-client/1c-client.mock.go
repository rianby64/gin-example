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
func (mc *MockClient) GetEntryList() ([]interface{}, error) {
	return []interface{}{}, nil
}

// GetEntry for testing
func (mc *MockClient) GetEntry() (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ok",
	}, nil
}

// UpdateEntry for testing
func (mc *MockClient) UpdateEntry() (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ok",
	}, nil
}

// DeleteEntry for testing
func (mc *MockClient) DeleteEntry() (map[string]interface{}, error) {
	return map[string]interface{}{
		"status": "ok",
	}, nil
}
