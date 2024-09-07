// api.go

package restql

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// api defines the structure for an api object.
type api struct {
	Name        string `yaml:"name" json:"name"`
	Path        string `yaml:"path" json:"path"`
	Method      string `yaml:"method" json:"method"`
	Description string `yaml:"description" json:"description"`

	Query *struct {
		Schema string `yaml:"schema" json:"schema"`
	} `yaml:"query" json:"query"`

	Payload *struct {
		Schema string `yaml:"schema" json:"schema"`
	} `yaml:"payload" json:"payload"`

	DB *struct {
		Query string `yaml:"query" json:"query"`
	} `yaml:"db" json:"db"`
}

// apiService defines the interface for a service that can handle api requests.
type apiService interface {
	Handle(method, path string, handler gin.HandlerFunc) error
}

// generateAPI takes an api object containing properties such as path, method, etc., and an apiService object.
func generateAPI(api api, service apiService) error {
	method := api.Method
	path := api.Path

	if path == "" {
		return errors.New("missing path in api config")
	}
	if method == "" {
		return errors.New("missing method in api config")
	}
	return service.Handle(method, path, generateHandler(api))
}

// GenerateAPIs takes an array of api objects and an apiService object,
// and generates handlers for each api using the provided apiService.
func GenerateAPIs(apis []api, service apiService) error {
	for _, api := range apis {
		err := generateAPI(api, service)
		if err != nil {
			logError("Failed to generate api", api.Path, err.Error())
			return err
		}
		logInfo("Generated api", api.Path, api.Method)
	}
	return nil
}

func generateResponseData(jsonParams map[string]interface{}) gin.H {
	return gin.H{
		"params": jsonParams,
	}
}

func responseError(c *gin.Context, statusCode int, errors []ValidationError, params map[string]interface{}) {
	// Prepare the response data with errors as a slice of ValidationError
	responseData := generateResponseData(params)
	responseData["errors"] = errors

	c.JSON(statusCode, responseData)
}

func generateHandler(api api) gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := make(map[string]interface{})
		for key, values := range c.Request.URL.Query() {
			value := values[0]
			if intValue, err := strconv.Atoi(value); err == nil {
				queryParams[key] = intValue
			} else if floatValue, err := strconv.ParseFloat(value, 64); err == nil {
				queryParams[key] = floatValue
			} else if boolValue, err := strconv.ParseBool(value); err == nil {
				queryParams[key] = boolValue
			} else {
				queryParams[key] = value
			}
		}

		if api.Query != nil && api.Query.Schema != "" {
			errors := validateJSON(queryParams, api.Query.Schema)
			if len(errors) > 0 {
				responseError(c, http.StatusBadRequest, errors, queryParams)
				return
			}
		}

		var params map[string]interface{}
		if err := c.ShouldBindJSON(&params); len(params) > 0 && err != nil {
			responseError(c, http.StatusBadRequest, []ValidationError{{
				Key:     "body",
				Message: err.Error(),
			}}, params)
			return
		}

		if api.Payload != nil && api.Payload.Schema != "" {
			errors := validateJSON(params, api.Payload.Schema)
			if len(errors) > 0 {
				responseError(c, http.StatusBadRequest, errors, params)
				return
			}
		}

		if api.DB != nil {
			fmt.Println()
		}

		responseData := generateResponseData(params)
		c.JSON(http.StatusOK, responseData)
	}
}
