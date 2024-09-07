package restql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseQueryParams(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string][]string
		expected map[string]interface{}
	}{
		{
			name: "parse integers",
			input: map[string][]string{
				"age": {"25"},
			},
			expected: map[string]interface{}{
				"age": 25,
			},
		},
		{
			name: "parse floats",
			input: map[string][]string{
				"price": {"19.99"},
			},
			expected: map[string]interface{}{
				"price": 19.99,
			},
		},
		{
			name: "parse booleans",
			input: map[string][]string{
				"active": {"true"},
			},
			expected: map[string]interface{}{
				"active": true,
			},
		},
		{
			name: "parse strings",
			input: map[string][]string{
				"name": {"John"},
			},
			expected: map[string]interface{}{
				"name": "John",
			},
		},
		{
			name: "mixed types",
			input: map[string][]string{
				"id":     {"42"},
				"price":  {"99.99"},
				"valid":  {"false"},
				"status": {"active"},
			},
			expected: map[string]interface{}{
				"id":     42,
				"price":  99.99,
				"valid":  false,
				"status": "active",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseQueryParams(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"123", true},
		{"0", true},
		{"-123", true},
		{"123.45", false},
		{"abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			result := isInt(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsFloat(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"123.45", true},
		{"0", true},
		{"-123.45", true},
		{"123", true},
		{"abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			result := isFloat(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsBool(t *testing.T) {
	tests := []struct {
		value    string
		expected bool
	}{
		{"true", true},
		{"false", true},
		{"1", false},
		{"0", false},
		{"abc", false},
	}

	for _, tt := range tests {
		t.Run(tt.value, func(t *testing.T) {
			result := isBool(tt.value)
			assert.Equal(t, tt.expected, result)
		})
	}
}
