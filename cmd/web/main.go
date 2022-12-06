package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/louisgarwood/bookings/internal/config"
	"github.com/louisgarwood/bookings/internal/handlers"
	"github.com/louisgarwood/bookings/internal/models"
	"github.com/louisgarwood/bookings/internal/render"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("starting application on port %s \n", portNumber)

	server := &http.Server{
		Addr:    portNumber,
		Handler: DefineRoutes(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() error {

	//what will I store in the session
	gob.Register(models.ReservationData{})

	app.IsProduction = false

	session = scs.New()
	session.Lifetime = 1 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.IsProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = templateCache
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	

	return nil
}
