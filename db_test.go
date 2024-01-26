package restql

import (
	"os"
	"testing"
)

func TestParseQueryFile(t *testing.T) {
	path := "./test/data/input_template_test.sql.tpl"
	data := map[string]interface{}{
		"TableName": "example_table",
		"Condition": "column = 'value'",
	}

	expectedOutputFilePath := "./test/data/output_template_test.sql"
	expectedOutput, err := os.ReadFile(expectedOutputFilePath)
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}

	output, err := parseQueryFile(path, data)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if output != string(expectedOutput) {
		t.Errorf("Expected output: %s, Got: %s", expectedOutput, output)
	}
}
