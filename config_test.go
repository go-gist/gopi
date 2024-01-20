// config_test.go

package restql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	filename := "test_data/restql.yml"

	// Call the readConfig function
	config, err := ReadConfig(filename)

	fmt.Printf("Configuration:\n%+v\n", config)

	// Check if there's an error
	assert.NoError(t, err, "Expected no error while reading the configuration file")

	// Check overall structure
	assert.NotNil(t, config, "Expected non-nil configuration")
	assert.NotNil(t, config.ParameterTemplates, "Expected non-nil ParameterTemplates")
	assert.NotNil(t, config.APIs, "Expected non-nil APIs")

	// Check specific values
	assert.Equal(t, "id", config.ParameterTemplates["common_id"].Name, "Unexpected common_id name")
	assert.Equal(t, "integer", config.ParameterTemplates["common_id"].Type, "Unexpected common_id type")
	assert.Equal(t, "Item id", config.ParameterTemplates["common_id"].Description, "Unexpected common_id description")

	// Check the first API
	assert.Equal(t, "Foo Get", config.APIs[0].Name, "Unexpected API name")
	assert.Equal(t, "/foo", config.APIs[0].Path, "Unexpected API path")
	assert.Equal(t, "GET", config.APIs[0].Method, "Unexpected API method")
	assert.Equal(t, "Get an item", config.APIs[0].Description, "Unexpected API description")
	assert.Len(t, config.APIs[0].Parameters, 1, "Unexpected number of parameters in API")
	assert.Equal(t, "id", config.APIs[0].Parameters[0].Name, "Unexpected parameter name in API")

	// Check the second API
	assert.Equal(t, "Foo Set", config.APIs[1].Name, "Unexpected API name")
	assert.Equal(t, "/foo", config.APIs[1].Path, "Unexpected API path")
	assert.Equal(t, "POST", config.APIs[1].Method, "Unexpected API method")
	assert.Equal(t, "Creates a foo item", config.APIs[1].Description, "Unexpected API description")
	assert.Len(t, config.APIs[1].Parameters, 2, "Unexpected number of parameters in API")
	assert.Equal(t, "id", config.APIs[1].Parameters[0].Name, "Unexpected parameter name in API")
	assert.Equal(t, "name", config.APIs[1].Parameters[1].Name, "Unexpected parameter name in API")
	assert.Equal(t, "string", config.APIs[1].Parameters[1].Type, "Unexpected parameter type in API")
	assert.Equal(t, "Item name", config.APIs[1].Parameters[1].Description, "Unexpected parameter description in API")
}
