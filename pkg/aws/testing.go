package aws

import "context"

// MockKMSClient is a mock implementation of KMSClient.
type MockKMSClient struct {
	// DataKey is the data key used for testing. It is exposed,
	// in case the test needs to refer/overwrite it.
	DataKey []byte
}

// NewMockKMSClient creates a new MockClient with a default data key.
func NewMockKMSClient() *MockKMSClient {
	return &MockKMSClient{
		DataKey: []byte{189, 210, 91, 253, 139, 50, 99, 24, 216, 87, 255, 114, 240, 71, 14, 182, 104, 251, 46, 62, 206, 22, 229, 117, 121, 102, 115, 81, 72, 99, 43, 194},
	}
}

// CreateDataKey implements a DataKeyManagementClient method.
func (m *MockKMSClient) CreateDataKey(ctx context.Context) ([]byte, []byte, error) {
	return m.DataKey, m.DataKey, nil
}

// DecryptDataKey implements a DataKeyManagementClient method.
func (m *MockKMSClient) DecryptDataKey(ctx context.Context, encryptedKey []byte) ([]byte, error) {
	return m.DataKey, nil
}
