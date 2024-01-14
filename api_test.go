// restql_test.go

package restql

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockAPIService is a mock implementation of the APIService interface for testing purposes.
type MockAPIService struct {
	HandledMethod string
	HandledPath   string
}

// Handle is the implementation of the APIService interface for testing purposes.
func (m *MockAPIService) Handle(method, path string, handler gin.HandlerFunc) {
	m.HandledMethod = method
	m.HandledPath = path
}

func TestGenerateAPI(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}

	// Create an instance of the API object for testing
	testAPI := API{
		path:   "/test",
		method: "POST",
	}

	// Call GenerateAPI with the mock service
	GenerateAPI(testAPI, mockService)

	// Assert that the Handle method of the mock service was called with the correct parameters
	assert.Equal(t, "POST", mockService.HandledMethod)
	assert.Equal(t, "/test", mockService.HandledPath)
}
