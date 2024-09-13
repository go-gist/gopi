// api_gin_test.go

package rest

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandle(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	ginAPI := GinAPIService{Engine: gin.New()}

	testHandler := func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Test Handler"})
	}

	supportedMethods := []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS", "HEAD"}
	for _, method := range supportedMethods {
		err := ginAPI.Handle(method, "/test", testHandler)

		if err != nil {
			t.Errorf("Expected no error for method %s, got: %v", method, err)
		}

		req, _ := http.NewRequest(method, "/test", nil)
		resp := httptest.NewRecorder()
		ginAPI.Engine.ServeHTTP(resp, req)

		if resp.Code != http.StatusOK {
			t.Errorf("Expected status code %d for method %s, got: %d", http.StatusOK, method, resp.Code)
		}

		expectedBody := `{"message":"Test Handler"}`
		if resp.Body.String() != expectedBody {
			t.Errorf("Expected response body %s for method %s, got: %s", expectedBody, method, resp.Body.String())
		}
	}

	unsupportedMethod := "INVALID"
	err := ginAPI.Handle(unsupportedMethod, "/test", testHandler)

	expectedError := errors.New("unsupported HTTP method")
	if err == nil || err.Error() != expectedError.Error() {
		t.Errorf("Expected error %v for unsupported method, got: %v", expectedError, err)
	}
}
