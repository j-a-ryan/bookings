package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Avoids config cycles that destroy compilability

// AppConfig holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
