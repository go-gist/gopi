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

// API defines the structure for an API object.
type API struct {
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

// APIService defines the interface for a service that can handle API requests.
type APIService interface {
	Handle(method, path string, handler gin.HandlerFunc) error
}

// GenerateAPI takes an API object containing properties such as path, method, etc., and an APIService object.
func GenerateAPI(api API, service APIService) error {
	method := api.Method
	path := api.Path

	if path == "" {
		return errors.New("missing path in API config")
	}
	if method == "" {
		return errors.New("missing method in API config")
	}
	return service.Handle(method, path, generateHandler(api))
}

// GenerateAPIs takes an array of API objects and an APIService object,
// and generates handlers for each API using the provided APIService.
func GenerateAPIs(apis []API, service APIService) error {
	for _, api := range apis {
		err := GenerateAPI(api, service)
		if err != nil {
			LogError("Failed to generate API", api.Path, err.Error())
			return err
		}
		LogInfo("Generated API", api.Path, api.Method)
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

func generateHandler(api API) gin.HandlerFunc {
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
	schemaContent, err := os.ReadFile(ConfigPath + "/" + schemaFilePath)
	if err != nil {
		Log.Error("Missing JSON schema", err)
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
