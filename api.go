// api.go

package restql

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/xeipuuv/gojsonschema"
)

// api defines the structure for an api object.
type api struct {
	Name        string `yaml:"name" json:"name"`
	Path        string `yaml:"path" json:"path"`
	Method      string `yaml:"method" json:"method"`
	Description string `yaml:"description" json:"description"`

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
		"json": jsonParams,
	}
}

func responseError(c *gin.Context, statusCode int, errorMessage string, params map[string]interface{}) {
	responseData := generateResponseData(params)
	responseData["error"] = errorMessage

	c.JSON(statusCode, responseData)
}

func generateHandler(api api) gin.HandlerFunc {
	return func(c *gin.Context) {
		var params map[string]interface{}

		if err := c.ShouldBindJSON(&params); err != nil {
			responseError(c, http.StatusBadRequest, err.Error(), params)
			return
		}

		// Validate JSON payload against the schema
		err := validateJSON(params, api.Payload.Schema)
		if err != nil {
			responseError(c, http.StatusBadRequest, err.Error(), params)
			return
		}

		if api.DB != nil {
			fmt.Println()
		}

		responseData := generateResponseData(params)
		c.JSON(http.StatusOK, responseData)
	}
}

func validateJSON(data map[string]interface{}, schemaFilePath string) error {
	// Read the JSON schema from the file
	schemaContent, err := os.ReadFile(configBasePath + "/" + schemaFilePath)
	if err != nil {
		logError("Missing JSON schema", err)
		return err
	}

	loader := gojsonschema.NewStringLoader(string(schemaContent))
	document := gojsonschema.NewGoLoader(data)

	result, err := gojsonschema.Validate(loader, document)
	if err != nil {
		return err
	}

	if !result.Valid() {
		// Handle validation errors
		var errors []string
		for _, err := range result.Errors() {
			errors = append(errors, err.String())
		}
		return fmt.Errorf("JSON schema validation failed: %v", errors)
	}

	return nil
}
