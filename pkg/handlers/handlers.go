package handlers

import (
	"github.com/louisgarwood/bookings/pkg/config"
	"github.com/louisgarwood/bookings/pkg/models"
	"github.com/louisgarwood/bookings/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (repo *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	repo.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {

	remoteIP := repo.App.Session.GetString(request.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["message"] = "Well hello there."
	stringMap["remote_ip"] = remoteIP
	
	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (repo *Repository) Lou(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "lou.page.tmpl", &models.TemplateData{})
}
