package restql

func parseQueryFile(path string, data map[string]interface{}) (string, error) {
	templateFromFile, err := LoadTemplateFromFile(path)
	if err != nil {
		return "", err
	}

	output, err := ExecuteTemplate(templateFromFile, data)
	if err != nil {
		return "", err
	}

	return output, nil
}
