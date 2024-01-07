// parse_template_test.go

package gopi

import (
	"testing"
)

func TestParseTemplate(t *testing.T) {
	data := struct {
		Name string
	}{
		Name: "cosmos",
	}

	tmpl := "Hello, {{.Name}}!"

	result, err := ParseTemplate(data, tmpl)
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}

	expected := "Hello, cosmos!"
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
