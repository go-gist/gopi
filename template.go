package restql

import (
	"bytes"
	"os"
	"text/template"
)

// loadTemplateFromFile reads a file and returns a parsed template based on the file content.
// It takes the file path as input and returns a template or an error if the file cannot be read or parsed.
func loadTemplateFromFile(filePath string) (*template.Template, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return template.New(filePath).Parse(string(content))
}

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
