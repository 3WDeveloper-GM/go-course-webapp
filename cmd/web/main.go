// A little web application
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/3WDeveloper-GM/go-course-webapp/pkg/config"
	"github.com/3WDeveloper-GM/go-course-webapp/pkg/handlers"
	"github.com/3WDeveloper-GM/go-course-webapp/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const portnumber = ":9060" //The port number for the application

var app config.AppConfig        //I'm using this variable to make a web-app wide configuration across all the files.
var session *scs.SessionManager // This variable stores the sessions.

func main() {

	// This variable is made in order to make the configuration of my app accesible to most of the methods in my program
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction //In production this is true.

	app.Session = session

	// I create a template cache here in order to cache the templates of my web app, this is done in order to load layouts and the html of my webpages separately
	tc, err := render.CreateTemplateCache()
	if err != nil { //standard golang error handling
		// stop the app if the template cache cannot be created.
		log.Fatal("cannot create template cache", err)
	}

	// I assign the template cache in the tc variable (accesible to handlers and render packages)
	app.TemplateCache = tc

	// This value is assigned in order to use the cached version of my webfiles (faster, but i cannot edit the html files in order to hot reload the websites, or creating a new version of the cache in order to reload the websites every time I reload the web page.)
	app.UseCache = false

	repo := handlers.NewRepo(&app) // I think this line makes the handlers.go methods accesible to the main.go file.

	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portnumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
