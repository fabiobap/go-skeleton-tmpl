package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/fabiobap/go-pdf-optimizer/helpers"
	"github.com/fabiobap/go-pdf-optimizer/internal/config"
	"github.com/fabiobap/go-pdf-optimizer/internal/handlers"
	"github.com/fabiobap/go-pdf-optimizer/internal/render"
)

const HTTP_PORT = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    HTTP_PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(map[string]int{})

	app.InProduction = false
	app.UseCache = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
		return err
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return nil
}
