package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/models"
)

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(templateData *models.TemplateData, request *http.Request) *models.TemplateData {
	templateData.CSRFToken = nosurf.Token(request)
	return templateData
}

func RenderTemplate(writer http.ResponseWriter, tmpl string, templateData *models.TemplateData, request *http.Request) {

	var templateCache map[string]*template.Template
	var err error

	if app.UseCache {
		templateCache = app.TemplateCache
	} else {
		templateCache, err = CreateTemplateCache()
		if err != nil {
			log.Fatal("could not create template cache")
		}
	}

	//get requested template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("temlate not found in cache")
	}

	buf := new(bytes.Buffer)

	templateData = AddDefaultData(templateData, request)

	err = template.Execute(buf, templateData)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(writer)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {

	log.Println("creating template cache...")
	// var templateCache = make(map[string]*template.Template)
	templateCache := map[string]*template.Template{}

	//get all of the files named *.page.tmpl
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}

	//get all of the files named *.layout.tmpl
	layouts, err := filepath.Glob("./templates/*.layout.tmpl")
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
			templateSet, err = templateSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[fileName] = templateSet
	}
	return templateCache, nil
}
