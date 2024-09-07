package restql

import (
	"os"
	"testing"
)

func TestLoadTemplateFromFile(t *testing.T) {
	// Path to the existing template file
	templateFilePath := "./test-config/foo.sql.tpl"

	// Test loading the template from the existing file
	templateFromFile, err := loadTemplateFromFile(templateFilePath)
	if err != nil {
		t.Fatalf("Failed to load template: %v", err)
	}

	// Ensure the template is not nil
	if templateFromFile == nil {
		t.Fatal("Template is nil")
	}

	// Define data to be used by the template
	data := map[string]interface{}{
		"TableName": "users",
		"Start":     0,
		"Size":      2,
	}
	// Execute the template with the provided data
	output, err := executeTemplate(templateFromFile, data)
	if err != nil {
		t.Fatalf("Failed to execute template: %v", err)
	}

	// Read the expected output from the file
	expectedOutputFilePath := "./test-config/foo.sql"
	expectedOutput, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}

	// Ensure the output matches the expected output
	if string(output) != string(expectedOutput) {
		t.Fatalf("Unexpected output. Expected: %s, Got: %s", string(expectedOutput), output)
	}
}
