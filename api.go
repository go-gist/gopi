// api.go

package restql

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// APIService defines the interface for a service that can handle API requests.
type APIService interface {
	Handle(method, path string, handler gin.HandlerFunc) error
}

// GenerateAPI takes an API object containing properties such as path, method, etc., and an APIService object.
func GenerateAPI(api API, service APIService) error {
	method := api.Method
	path := api.Path

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

func validateParameters(c *gin.Context, params []APIParameter) error {
	// Validate parameters
	for _, paramSpec := range params {
		// Check if the parameter is provided
		paramValue := c.Query(paramSpec.Name)

		if paramSpec.Required && paramValue == "" {
			return errors.New("'" + paramSpec.Name + "' is a required parameter")
		}

		// Validate based on the type only if the parameter is provided
		if paramValue != "" {
			switch paramSpec.Type {
			case "integer":
				_, err := strconv.Atoi(paramValue)
				if err != nil {
					return errors.New("'" + paramSpec.Name + "' must be an integer")
				}
			case "string":
				// Additional string validation logic if needed
			case "number":
				_, err := strconv.ParseFloat(paramValue, 64)
				if err != nil {
					return errors.New("'" + paramSpec.Name + "' must be a number")
				}
			case "boolean":
				_, err := strconv.ParseBool(paramValue)
				if err != nil {
					return errors.New("'" + paramSpec.Name + "' must be a boolean")
				}
			default:
				return errors.New("Unsupported type: '" + paramSpec.Type + "'")
			}
		}
	}

	return nil
}

func generateHandler(api API) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Validate parameters
		if err := validateParameters(c, api.Parameters); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// If all validations pass, you can proceed with handling the API logic
		c.JSON(http.StatusOK, api)
	}
}
