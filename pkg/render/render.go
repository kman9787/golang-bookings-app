package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/kman9787/golang-bookings-app/pkg/config"
	"github.com/kman9787/golang-bookings-app/pkg/models"
)

const templateDir = "./templates"

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var templateCache map[string]*template.Template

	if app.UseCache {
		// Crate temlate cache
		templateCache = app.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	// Get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Unable to retrieve template cache")
	}

	td = AddDefaultData(td)
	buf := new(bytes.Buffer)
	_ = template.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("Failed to write template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	pages, err := filepath.Glob(templateDir + "/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		layouts, err := filepath.Glob(templateDir + "/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseFiles(layouts...)
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[name] = templateSet
	}

	return templateCache, nil
}
