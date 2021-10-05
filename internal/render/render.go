package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/leetrent/bookings/internal/config"
	"github.com/leetrent/bookings/internal/models"
)

var functions = template.FuncMap{}

var appConfig *config.AppConfig

func NewTemplates(ac *config.AppConfig) {
	appConfig = ac
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = appConfig.Session.PopString(r.Context(), "flash")
	td.Error = appConfig.Session.PopString(r.Context(), "error")
	td.Warning = appConfig.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(responseWriter http.ResponseWriter, r *http.Request, templateName string, templateData *models.TemplateData) {
	var templateCache map[string]*template.Template

	if appConfig.UseCache {
		templateCache = appConfig.TemplateCache
	} else {
		templateCache, _ = CreateTemplateCache()
	}

	template, ok := templateCache[templateName]
	if !ok {
		log.Fatal("Error encountered whan attempting to get template from template cache for ", templateName)
	}

	buffer := new(bytes.Buffer)
	templateData = AddDefaultData(templateData, r)
	_ = template.Execute(buffer, templateData)

	//nbrOfBytes, err := buffer.WriteTo(responseWriter)
	_, err := buffer.WriteTo(responseWriter)
	if err != nil {
		fmt.Println("Error writing template to response writer", err)
	}

	//fmt.Println("# of bytes written:", nbrOfBytes)
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		// fmt.Println("Page is currently", page)
		// fmt.Println("Name is currently", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//matches, err := filepath.Glob("./templates/*.layout.tmpl")
		matches, err := filepath.Glob(".\\templates\\*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			//ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			ts, err = ts.ParseGlob(".\\templates\\*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
