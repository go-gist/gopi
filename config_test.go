// config_test.go

package restql

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const filename = "test_data/restql.yml"

func TestReadConfig(t *testing.T) {
	config, err := GetAPIConfig(filename)

	assert.NoError(t, err, "Expected no error while reading the configuration file")

	assert.NotNil(t, config, "Expected non-nil configuration")
	assert.NotNil(t, config.APIs, "Expected non-nil APIs")

	assert.Equal(t, "Foo Get", config.APIs[0].Name, "Unexpected API name")
	assert.Equal(t, "/foo", config.APIs[0].Path, "Unexpected API path")
	assert.Equal(t, "GET", config.APIs[0].Method, "Unexpected API method")
	assert.Equal(t, "Get an item", config.APIs[0].Description, "Unexpected API description")
	assert.Len(t, config.APIs[0].Parameters, 1, "Unexpected number of parameters in API")
	assert.Equal(t, "id", config.APIs[0].Parameters[0].Name, "Unexpected parameter name in API")
	assert.Equal(t, true, config.APIs[0].Parameters[0].Required, "Unexpected parameter name in API")

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

func TestGetAPIs(t *testing.T) {
	config, _ := GetAPIConfig(filename)
	apis := GetAPIs(config)

	assert.Len(t, apis, len(config.APIs), "Unexpected number of APIs")
}
