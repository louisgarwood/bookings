package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/models"
	"github.com/louisgarwood/bookings/internal/render"
)

func TestMain(m *testing.M) {
	getRoutes()
}

func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

var app config.AppConfig
var session *scs.SessionManager

func getRoutes() http.Handler {
	//what will I store in the session
	gob.Register(models.ReservationData{})

	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	app.Session = session

	templateCache, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = true
	render.NewTemplates(&app)

	repo := NewRepo(&app)
	NewHandlers(repo)

	multiplexer := chi.NewRouter()

	multiplexer.Use(middleware.Recoverer)
	multiplexer.Use(SessionLoad)

	multiplexer.Get("/", Repo.Home)
	multiplexer.Get("/about", Repo.About)
	multiplexer.Get("/contact", Repo.Contact)
	multiplexer.Get("/captains", Repo.Captains)
	multiplexer.Get("/oceanview", Repo.Oceanview)

	multiplexer.Get("/search-availability", Repo.SearchAvailability)
	multiplexer.Post("/search-availability", Repo.PostSearchAvailability)
	multiplexer.Post("/search-availability-json", Repo.AvailabilityJSON)

	multiplexer.Get("/make-reservation", Repo.Reservation)
	multiplexer.Post("/make-reservation", Repo.PostReservation)

	multiplexer.Get("/reservation-summary", Repo.ReservationSummary)

	fileServer := http.FileServer(http.Dir("./static/"))
	multiplexer.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return multiplexer
}

var pathToTemplates string = "../../templates"

func CreateTestTemplateCache() (map[string]*template.Template, error) {

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
