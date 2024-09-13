package rest

import (
	"bytes"
	"text/template"
)

// executeTemplate executes the provided template with the given data and returns the result as a string.
// It uses a buffer to capture the output and returns an error if template execution fails.
func executeTemplate(tmpl *template.Template, data interface{}) (string, error) {
	var result string
	// Use a buffer to capture the output
	outputBuffer := &bytes.Buffer{}
	err := tmpl.Execute(outputBuffer, data)
	if err != nil {
		return "", err
	}

	// Convert the buffer to a string
	result = outputBuffer.String()
	return result, nil
}
