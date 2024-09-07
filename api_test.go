// api_test.go

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
func (m *MockAPIService) Handle(method, path string, handler gin.HandlerFunc) error {
	m.HandledMethod = method
	m.HandledPath = path
	return nil
}

func TestGenerateAPI(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}

	// Create an instance of the API object for testing
	testAPI := API{
		Path:   "/foo",
		Method: "POST",
	}

	// Call GenerateAPI with the mock service
	err := GenerateAPI(testAPI, mockService)

	// Assert that the Handle method of the mock service was called with the correct parameters
	assert.Equal(t, "POST", mockService.HandledMethod)
	assert.Equal(t, "/foo", mockService.HandledPath)
	assert.NoError(t, err)
}

func TestGenerateAPIs_Error(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}

	// Create an instance of the API object for testing with missing fields
	testAPI := API{}

	// Call GenerateAPIs with the mock service and an API with missing fields
	err := GenerateAPIs([]API{testAPI}, mockService)

	// Assert that an error is returned
	assert.Error(t, err)
}

func TestGenerateAPIs_Success(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}

	// Define a valid list of API objects for testing
	validAPIs := []API{
		{
			Path:   "/foo",
			Method: "GET",
		},
		{
			Path:   "/bar",
			Method: "POST",
		},
	}

	// Call GenerateAPIs with the mock service and a list of valid APIs
	err := GenerateAPIs(validAPIs, mockService)

	// Assert that the Handle method of the mock service was called with the correct parameters
	// We can verify this by checking the last set of parameters that were handled
	if len(validAPIs) > 0 {
		lastAPI := validAPIs[len(validAPIs)-1]
		assert.Equal(t, lastAPI.Method, mockService.HandledMethod)
		assert.Equal(t, lastAPI.Path, mockService.HandledPath)
	}

	// Assert that no error is returned
	assert.NoError(t, err)
}
