// restql_test.go

package restql

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGinAPIService_Handle(t *testing.T) {
	// Set Gin to release mode during testing
	gin.SetMode(gin.ReleaseMode)

	tests := []struct {
		name   string
		method string
		path   string
	}{
		{name: "Test GET request", method: "GET", path: "/test-get"},
		{name: "Test POST request", method: "POST", path: "/test-post"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new Gin engine for testing
			engine := gin.New()

			// Create a GinAPIService instance
			apiService := &GinAPIService{
				Engine: engine,
			}

			// Define a sample handler function for testing
			sampleHandler := func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{"message": "success"})
			}

			// Test handling request
			apiService.Handle(tt.method, tt.path, sampleHandler)

			// Create a test request
			req, err := http.NewRequest(tt.method, tt.path, nil)
			assert.NoError(t, err)

			// Create a response recorder
			w := httptest.NewRecorder()

			// Perform the request
			engine.ServeHTTP(w, req)

			// Assert the status code and response body
			assert.Equal(t, http.StatusOK, w.Code)
			assert.Contains(t, w.Body.String(), "success")
		})
	}
}
