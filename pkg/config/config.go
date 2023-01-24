package config

import (
	"html/template"
	"log"
)

// AppConfig holds the application config
type AppConfig struct {
	UseCatche     bool
	TamplateCatch map[string]*template.Template
	InfoLog       *log.Logger
}
