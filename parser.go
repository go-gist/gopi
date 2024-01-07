// parse_template.go

package gopi

import (
	"bytes"
	"text/template"
)

// ParseTemplate parses a template and returns the result.
func ParseTemplate(data interface{}, tmpl string) (string, error) {
	t, err := template.New("parseTemplate").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var result bytes.Buffer
	err = t.Execute(&result, data)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
