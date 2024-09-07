package restql

func parseQueryFile(path string, data map[string]interface{}) (string, error) {
	templateFromFile, err := loadTemplateFromFile(path)
	if err != nil {
		logError("Failed to load DB query file", err.Error())
		return "", err
	}

	output, err := executeTemplate(templateFromFile, data)
	if err != nil {
		return "", err
	}

	return output, nil
}
