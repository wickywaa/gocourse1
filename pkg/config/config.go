package config

import "html/template"

type AppConfig struct {

	// AppConfig holds the application config
	UseCache      bool
	TemplateCache map[string]*template.Template
}
