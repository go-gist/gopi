package rest

import (
	"fmt"
	"os"

	"github.com/xeipuuv/gojsonschema"
)

type ValidationError struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

func validateJSON(data map[string]interface{}, schemaFilePath string) []ValidationError {
	var validationErrors []ValidationError

	// Construct the full path to the schema file
	fullPath := fmt.Sprintf("%s%s", configBasePath, schemaFilePath)

	// Read the schema file content
	schemaContent, err := os.ReadFile(fullPath)
	if err != nil {
		validationErrors = append(validationErrors, ValidationError{
			Key:     "schema",
			Message: fmt.Sprintf("error reading schema file %s: %v", fullPath, err),
		})
		return validationErrors
	}

	// Create JSON schema loader from schema content
	loader := gojsonschema.NewStringLoader(string(schemaContent))
	document := gojsonschema.NewGoLoader(data)

	// Validate the JSON document against the schema
	result, err := gojsonschema.Validate(loader, document)
	if err != nil {
		validationErrors = append(validationErrors, ValidationError{
			Key:     "validation",
			Message: fmt.Sprintf("error validating JSON schema: %v", err),
		})
		return validationErrors
	}

	// Collect validation errors if the document is invalid
	if !result.Valid() {
		for _, err := range result.Errors() {
			validationErrors = append(validationErrors, ValidationError{
				Key:     err.Field(),
				Message: err.String(),
			})
		}
	}

	return validationErrors
}
