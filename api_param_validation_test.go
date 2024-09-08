package restql

import (
	"testing"
)

func TestValidateJSON_ValidData(t *testing.T) {
	params := map[string]interface{}{
		"size":  10,
		"start": 2,
	}

	schemaPath := "test-config/foo_query_schema.json"
	errors := validateJSON(params, schemaPath)

	if len(errors) > 0 {
		t.Errorf("expected no errors, got %v", errors)
	}
}

func TestValidateJSON_MissingRequiredFields(t *testing.T) {
	params := map[string]interface{}{
		"size": 10, // Missing "start"
	}

	schemaPath := "./test-config/foo_query_schema.json"
	errors := validateJSON(params, schemaPath)

	expectedErrors := []ValidationError{
		{Key: "(root)", Message: "(root): start is required"},
	}

	if len(errors) == 0 {
		t.Fatal("expected errors, got none")
	}

	if len(errors) != len(expectedErrors) {
		t.Fatalf("expected %d error(s), got %d", len(expectedErrors), len(errors))
	}

	for i, err := range errors {
		if err != expectedErrors[i] {
			t.Errorf("expected error %v, got %v", expectedErrors[i], err)
		}
	}
}

func TestValidateJSON_InvalidFieldTypes(t *testing.T) {
	params := map[string]interface{}{
		"size":  "ten", // Invalid type
		"start": 2,
	}

	schemaPath := "test-config/foo_query_schema.json"
	errors := validateJSON(params, schemaPath)

	expectedErrors := []ValidationError{
		{Key: "size", Message: "size: Invalid type. Expected: number, given: string"},
	}

	if len(errors) == 0 {
		t.Fatal("expected errors, got none")
	}

	if len(errors) != len(expectedErrors) {
		t.Fatalf("expected %d error(s), got %d", len(expectedErrors), len(errors))
	}

	for i, err := range errors {
		if err != expectedErrors[i] {
			t.Errorf("expected error %v, got %v", expectedErrors[i], err)
		}
	}
}

func TestValidateJSON_FieldValuesOutOfRange(t *testing.T) {
	params := map[string]interface{}{
		"size":  1001, // Out of range
		"start": 2,
	}

	schemaPath := "./test-config/foo_query_schema.json"
	errors := validateJSON(params, schemaPath)

	expectedErrors := []ValidationError{
		{Key: "size", Message: "size: Must be less than or equal to 1000"},
	}

	if len(errors) == 0 {
		t.Fatal("expected errors, got none")
	}

	if len(errors) != len(expectedErrors) {
		t.Fatalf("expected %d error(s), got %d", len(expectedErrors), len(errors))
	}

	for i, err := range errors {
		if err != expectedErrors[i] {
			t.Errorf("expected error %v, got %v", expectedErrors[i], err)
		}
	}
}

func TestValidateJSON_ExtraFields(t *testing.T) {
	params := map[string]interface{}{
		"size":  10,
		"start": 2,
		"extra": "value", // Extra field
	}

	schemaPath := "./test-config/foo_query_schema.json"
	errors := validateJSON(params, schemaPath)

	if len(errors) > 0 {
		t.Errorf("expected no errors, got %v", errors)
	}
}
