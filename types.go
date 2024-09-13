package rest

import "text/template"

type apiConfig struct {
	APIs []api `yaml:"apis"`
}

type api struct {
	Name        string `yaml:"name" json:"name"`
	Path        string `yaml:"path" json:"path"`
	Method      string `yaml:"method" json:"method"`
	Description string `yaml:"description" json:"description"`

	Query *struct {
		Schema string `yaml:"schema" json:"schema"`
	} `yaml:"query" json:"query"`

	Payload *struct {
		Schema string `yaml:"schema" json:"schema"`
	} `yaml:"payload" json:"payload"`

	DB *struct {
		Query string `yaml:"query" json:"query"`
	} `yaml:"db" json:"db"`

	Actions []action `yaml:"actions" json:"actions"`
}

type action struct {
	Type          string `yaml:"type" json:"type"`
	Query         string `yaml:"query" json:"query"`
	queryTemplate *template.Template
}
