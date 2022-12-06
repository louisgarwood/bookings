package render

import (
	"net/http"
	"testing"

	"github.com/louisgarwood/bookings/internal/models"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	session.Put(request.Context(), "flash", "123")
	result := AddDefaultData(&td, request)
	if result.Flash != "123" {
		t.Error("Flash value of 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	request, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}

	context := request.Context()
	context, _ = session.Load(context, request.Header.Get("x-Session"))
	request = request.WithContext(context)

	return request, nil
}

func TestRenderTemplate(t *testing.T) {
	pathToTemplates = "./../../templates"
	templateCache, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = templateCache

	request, err := getSession()
	if err != nil {
		t.Error(err)
	}

	var ww myWriter

	err = RenderTemplate(&ww, "home.page.tmpl", &models.TemplateData{}, request)
	if err != nil {
		t.Error("error writing template to browser")
	}

	err = RenderTemplate(&ww, "nonexistent.page.tmpl", &models.TemplateData{}, request)
	if err == nil {
		t.Error("error writing template to browser")
	}
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
