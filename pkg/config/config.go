package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)


type AppConfig struct {
	TemplateCache map[string]*template.Template
	UseCache bool
	InfoLog *log.Logger
	Inproduction bool
	Session *scs.SessionManager
}