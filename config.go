// config.go

package rest

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// apiConfig represents the structure of the YAML configuration file

var configBasePath string

// readConfig reads the YAML configuration file
func GetAPIConfig(filename string) (*apiConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		logError("Configuration read failed", err.Error())
		return nil, err
	}
	configBasePath = fmt.Sprintf("%s/", filepath.Dir(filename))

	config := &apiConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		logError("Configuration parsing failed", err.Error())
		return nil, err
	}

	logInfo("Configuration parsing completed", filename)

	return config, nil
}

func GetAPIs(config *apiConfig) []api {
	var apis []api

	for _, apiEntry := range config.APIs {
		api := api(apiEntry)
		apis = append(apis, api)
	}

	return apis
}

func getConfigFullPath(filename string) string {
	return fmt.Sprintf("%s%s", configBasePath, filename)
}
