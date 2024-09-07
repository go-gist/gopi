// config_test.go

package restql

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const filename = "./test-config/restql.yml"

func TestReadConfig(t *testing.T) {
	// Define the expected values
	expectedAPIs := []struct {
		Name        string
		Path        string
		Method      string
		Description string
	}{
		{"Foo", "/foo", "GET", "Get an item"},
		{"Bar", "/bar", "POST", "Creates an item"},
	}

	// Read the configuration
	config, err := GetAPIConfig(filename)

	// Assert no error is returned
	assert.NoError(t, err, "Expected no error while reading the configuration file")

	// Assert the configuration is not nil
	assert.NotNil(t, config, "Expected non-nil configuration")
	assert.NotNil(t, config.APIs, "Expected non-nil APIs")

	// Assert the correct number of APIs are loaded
	assert.Len(t, config.APIs, len(expectedAPIs), "Unexpected number of APIs")

	// Check each api against the expected values
	for i, expectedAPI := range expectedAPIs {
		api := config.APIs[i]

		t.Run(fmt.Sprintf("api #%d", i), func(t *testing.T) {
			assert.Equal(t, expectedAPI.Name, api.Name, "Unexpected api name")
			assert.Equal(t, expectedAPI.Path, api.Path, "Unexpected api path")
			assert.Equal(t, expectedAPI.Method, api.Method, "Unexpected api method")
			assert.Equal(t, expectedAPI.Description, api.Description, "Unexpected api description")
		})
	}
}

func TestGetAPIs(t *testing.T) {
	// Read the configuration
	config, err := GetAPIConfig(filename)
	assert.NoError(t, err, "Expected no error while reading the configuration file")
	assert.NotNil(t, config, "Expected non-nil configuration")
	assert.NotNil(t, config.APIs, "Expected non-nil APIs")

	// Get APIs
	apis := GetAPIs(config)

	// Assert the number of APIs matches the expected count
	assert.Len(t, apis, len(config.APIs), "Unexpected number of APIs")

	for i, api := range apis {
		t.Run(fmt.Sprintf("api #%d", i), func(t *testing.T) {
			assert.NotEmpty(t, api.Name, "api name should not be empty")
			assert.NotEmpty(t, api.Path, "api path should not be empty")
			assert.NotEmpty(t, api.Method, "api method should not be empty")
		})
	}
}
