package restql

import (
	"os"
	"testing"
)

func TestParseQueryFile(t *testing.T) {
	// Path to the SQL template
	path := "./test-config/foo.sql.tpl"

	// Data to be injected into the template
	data := map[string]interface{}{
		"TableName": "employees",
		"Filters": []string{
			"department = 'HR'",
			"hire_date >= '2023-01-01'",
		},
	}

	// Path to the file containing the expected output
	expectedOutputFilePath := "./test-config/foo.sql"

	// Read the expected output from file
	expectedOutput, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}

	// Parse the query file with the provided data
	output, err := parseQueryFile(path, data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Compare the actual output with the expected output
	if output != string(expectedOutput) {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, output)
	}
}
