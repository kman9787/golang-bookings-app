package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// Holds the application config
type AppConfig struct {
	PortNumber    string
	UseCache      bool
	TemplateCache map[string]*template.Template
	TemplateDir   string
	InfoLog       *log.Logger
	InProduction  bool
	Session       *scs.SessionManager
}
