package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/j-a-ryan/bookings/pkg/config"
	"github.com/j-a-ryan/bookings/pkg/models"
)

// RenderTemplate uses authomatic caching of all tmpl files in the templates
// folder.
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCacheAutomatic() // We're in dev mode, read new HTML every time it's called.
	}

	// Get requested template from cache
	t, ok := tc[tmpl] // ok true if found
	if !ok {
		log.Fatal("Could not get template from cache.")
	}

	buf := new(bytes.Buffer)
	td = AddDefaultData(td)
	_ = t.Execute(buf, td)

	// Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

var app *config.AppConfig

// NewTemplates set the config file for the template package and other config vars
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	// TODO: Put default data here
	return td
}

func CreateTemplateCacheAutomatic() (map[string]*template.Template, error) {
	// tmplCache := make(map[string]*template.Template) ... or same thing:
	tmplCache := map[string]*template.Template{}

	// Populate cache fully from templates folder. Parser requires the non-base
	// templates to be passed into it first, followed by base/dependency templates.
	// So, first get all *.page.tmpl file paths.
	pages, err := filepath.Glob("./templates/*.page.tmpl") // Makes pages a slice of strings
	if err != nil {
		return tmplCache, err
	}

	// Range through the files *.page
	for _, page := range pages {
		name := filepath.Base(page)                    // Returns the filename portion of the path
		ts, err := template.New(name).ParseFiles(page) // ts: template set
		if err != nil {
			return tmplCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl") // Get the dependency templates
		if err != nil {
			return tmplCache, err
		}
		if len(matches) > 0 { // If there are any layout files...
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl") // Add parsed layout.tmpls to ts set
			if err != nil {
				return tmplCache, err
			}
		}

		tmplCache[name] = ts // Add the resulting parsed template to the map
	}

	return tmplCache, nil
}
