package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/models"
	"github.com/louisgarwood/bookings/internal/render"
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

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{}, request)
}

func (repo *Repository) About(writer http.ResponseWriter, request *http.Request) {

	remoteIP := repo.App.Session.GetString(request.Context(), "remote_ip")
	stringMap := make(map[string]string)
	stringMap["message"] = "Well hello there."
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	}, request)
}

func (repo *Repository) Contact(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "contact.page.tmpl", &models.TemplateData{}, request)
}

func (repo *Repository) Captains(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "captains.page.tmpl", &models.TemplateData{}, request)
}

func (repo *Repository) Oceanview(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "oceanview.page.tmpl", &models.TemplateData{}, request)
}

func (repo *Repository) SearchAvailability(writer http.ResponseWriter, request *http.Request) {
	render.RenderTemplate(writer, "search-availability.page.tmpl", &models.TemplateData{}, request)
}

func (repo *Repository) PostSearchAvailability(writer http.ResponseWriter, request *http.Request) {
	start := request.Form.Get("start")
	end := request.Form.Get("end")

	writer.Write([]byte(fmt.Sprintf("Posted to search availability for dates: %s to %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (repo *Repository) AvailabilityJSON(writer http.ResponseWriter, request *http.Request) {
	response := jsonResponse{
		OK:      false,
		Message: "Available!",
	}

	output, error := json.MarshalIndent(response, "", "     ")
	if error != nil {
		log.Println(error)
	}

	writer.Header().Set("Content-type", "application/json")
	writer.Write(output)
}
