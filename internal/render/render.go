package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/models"
)

var app *config.AppConfig
var pathToTemplates string = "./templates"

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.Flash = app.Session.PopString(request.Context(), "flash")
	templateData.Warning = app.Session.PopString(request.Context(), "warning")
	templateData.Error = app.Session.PopString(request.Context(), "error")
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData, request *http.Request) error {

	var templateCache map[string]*template.Template
	var err error

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Println("could not create template cache")
			return errors.New("can't create template cache")
		}
	}

	//get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Println("template not found in cache")
		return errors.New("can't get template from cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, request)

	err = template.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
		return err
	}

	//render the template
	_, err = buf.WriteTo(writer)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	log.Println("creating template cache...")
	// var templateCache = make(map[string]*template.Template)
	templateCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	//get all of the files named *.layout.tmpl
	layouts, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
	if err != nil {
		return templateCache, err
	}

	//range through the pages
	for _, page := range pages {
		fileName := filepath.Base(page)
		templateSet, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		if len(layouts) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[fileName] = templateSet
	}
	return templateCache, nil
}
