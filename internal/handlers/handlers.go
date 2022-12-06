package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/forms"
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

func (repo *Repository) Reservation(writer http.ResponseWriter, request *http.Request) {
	var emptyReservation models.ReservationData

	data := make(map[string]interface{})
	data["reservation"] = emptyReservation

	render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
		Form: forms.New(nil),
		Data: data,
	}, request)
}

func (repo *Repository) PostReservation(writer http.ResponseWriter, request *http.Request) {
	error := request.ParseForm()
	if error != nil {
		log.Println(error)
		return
	}

	reservation := models.ReservationData{
		FirstName:    request.Form.Get("first-name"),
		LastName:     request.Form.Get("last-name"),
		EmailAddress: request.Form.Get("email-address"),
		PhoneNumber:  request.Form.Get("phone-number"),
	}

	form := forms.New(request.PostForm)

	form.Required("first-name", "last-name", "email-address")
	form.IsEmail("email-address")

	if !form.Valid() {
		data := make(map[string]interface{})
		data["reservation"] = reservation

		render.RenderTemplate(writer, "make-reservation.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		}, request)
		return
	}

	repo.App.Session.Put(request.Context(), "reservation", reservation)

	http.Redirect(writer, request, "/reservation-summary", http.StatusSeeOther)
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

func (repo *Repository) ReservationSummary(writer http.ResponseWriter, request *http.Request) {
	reservation, ok := repo.App.Session.Get(request.Context(), "reservation").(models.ReservationData)

	if !ok {
		log.Println("error getting reservation from session")
		repo.App.Session.Put(request.Context(), "error", "Cannot get reservation from sesson")
		http.Redirect(writer, request, "/", http.StatusTemporaryRedirect)
		return
	}

	repo.App.Session.Remove(request.Context(), "reservation")
	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.RenderTemplate(writer, "reservation-summary.page.tmpl", &models.TemplateData{Data: data}, request)
}
