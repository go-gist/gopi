// config.go

package restql

import (
	"os"

	"gopkg.in/yaml.v2"
)

// APIConfig represents the structure of the YAML configuration file
type APIConfig struct {
	APIs []API `yaml:"apis"`
}

// readConfig reads the YAML configuration file
func GetAPIConfig(filename string) (*APIConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &APIConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		Log.Error(err)
		return nil, err
	}

	return config, nil
}

func GetAPIs(config *APIConfig) []API {
	var apis []API

	for _, apiEntry := range config.APIs {
		api := API(apiEntry)
		apis = append(apis, api)
	}

	return apis
}
