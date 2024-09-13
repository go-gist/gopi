// config_test.go

package rest

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const filename = "./test-config/rest.yml"

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
