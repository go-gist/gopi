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
		LogError("Configuration read failed", err.Error())
		return nil, err
	}

	config := &APIConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		LogError("Configuration parsing failed", err.Error())
		return nil, err
	}

	LogInfo("Configuration parsing completed", filename)

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
