// api.go

package restql

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// APIService defines the interface for a service that can handle API requests.
type APIService interface {
	Handle(method, path string, handler gin.HandlerFunc) error
}

// GenerateAPI takes an API object containing properties such as path, method, etc., and an APIService object.
func GenerateAPI(api interface{}, service APIService) error {
	method := api.(API).Method
	path := api.(API).Path

	if method == "" {
		return errors.New("missing path in API object")
	}
	return service.Handle(method, path, generateHandler(api))
}

// GenerateAPIs takes an array of API objects and an APIService object,
// and generates handlers for each API using the provided APIService.
func GenerateAPIs(apis []API, service APIService) error {
	for _, api := range apis {
		err := GenerateAPI(api, service)
		if err != nil {
			return err
		}
	}
	return nil
}

// generateHandler is a helper function to generate the appropriate handler function based on the API object.
func generateHandler(api interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Handler for " + api.(API).Method + " " + api.(API).Path,
		})
	}
}
