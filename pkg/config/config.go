package config

import "html/template"

// AppConfig holds the application cofig
type AppConfig struct {
	TemplateCache map[string]*template.Template
}
