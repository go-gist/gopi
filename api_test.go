package rest

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// MockAPIService is a mock implementation of the apiService interface for testing purposes.
type MockAPIService struct {
	HandledMethod string
	HandledPath   string
}

// Handle is the implementation of the apiService interface for testing purposes.
func (m *MockAPIService) Handle(method, path string, handler gin.HandlerFunc) error {
	m.HandledMethod = method
	m.HandledPath = path
	return nil
}

func TestGenerateAPI(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}
	mockDbConnection := &SQL{}

	// Create an instance of the api object for testing
	testAPI := api{
		Path:   "/foo",
		Method: "POST",
	}

	// Call generateAPI with the mock service
	err := generateAPI(testAPI, mockService, mockDbConnection)

	// Assert that the Handle method of the mock service was called with the correct parameters
	assert.Equal(t, "POST", mockService.HandledMethod)
	assert.Equal(t, "/foo", mockService.HandledPath)
	assert.NoError(t, err)
}

func TestGenerateAPIs_Error(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &GinAPIService{}
	mockDbConnection := &SQL{}

	// Create an instance of the api object for testing with missing fields
	testAPI := api{}

	// Call GenerateAPIs with the mock service and an api with missing fields
	err := GenerateAPIs([]api{testAPI}, mockService, mockDbConnection)

	// Assert that an error is returned
	assert.Error(t, err)
}

func TestGenerateAPIs_Success(t *testing.T) {
	// Create an instance of MockAPIService
	mockService := &MockAPIService{}
	mockDbConnection := &SQL{}

	// Define a valid list of api objects for testing
	validAPIs := []api{
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
	err := GenerateAPIs(validAPIs, mockService, mockDbConnection)

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
