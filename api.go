package rest

import (
	"errors"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

// apiService defines the interface for a service that can handle api requests.
type apiService interface {
	Handle(method, path string, handler gin.HandlerFunc) error
}

// generateAPI takes an api object containing properties such as path, method, etc., and an apiService object.
func generateAPI(api api, service apiService, db dbConnection) error {
	method := api.Method
	path := api.Path

	if path == "" {
		return errors.New("missing path in api config")
	}
	if method == "" {
		return errors.New("missing method in api config")
	}

	for i, action := range api.Actions {
		if action.Type == "db" {
			content, err := os.ReadFile(getConfigFullPath(action.Query))
			if err != nil {
				return err
			}
			queryTemplate, err := template.New(action.Query).Parse(string(content))
			if err != nil {
				return err
			}
			api.Actions[i].queryTemplate = queryTemplate
		}
	}

	return service.Handle(method, path, generateHandler(api, db))
}

// GenerateAPIs takes an array of api objects and an apiService object,
// and generates handlers for each api using the provided apiService.
func GenerateAPIs(apis []api, service apiService, db dbConnection) error {
	for _, api := range apis {
		err := generateAPI(api, service, db)
		if err != nil {
			logError("Failed to generate api", api.Path, err.Error())
			return err
		}
		logInfo("Generated api", api.Path, api.Method)
	}
	return nil
}

func generateResponseData(jsonParams interface{}) gin.H {
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
