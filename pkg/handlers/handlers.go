package handlers

import (
	"fmt"
	"net/http"

	"github.com/3WDeveloper-GM/go-course-webapp/pkg/config"
	"github.com/3WDeveloper-GM/go-course-webapp/pkg/models"
	"github.com/3WDeveloper-GM/go-course-webapp/pkg/render"
)

// TemplateData holds data sent from handlers to handlers

// Repo is used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{App: a}
}

// Newhanlders sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// This method is responisble of handling the business logic that is performed every time the user enters in the home page of the project. In this case. This method can render the html via the render.RenderTemplate function and passes data around using the models.TemplateData struct.
func (repos *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIP := r.RemoteAddr
	repos.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	//n, err := fmt.Fprintf(w, "This is the home page")
	//if err != nil {
	//	log.Print(err)
	//}
	message :=
		`
		The number of bytes written
		when accesing the home page
		is equal to %v
		`
	fmt.Print(message)
}

// This method mostly performs the same logic as the one performed using the Home method. It executes business logic in the backend, and exposes it to the user using the models.TemplateData struct.
func (repos *Repository) About(w http.ResponseWriter, r *http.Request) {
	//n, err := fmt.Fprintf(w, "This is the about page, neat, yay, I can type moonspeak, check it out ,雪花")
	//if err != nil {
	//	log.Print(err)
	//}a

	//This is the business logic I mentioned, I create a "test" key and assign the value of "hello again" to it, this data is passed down to the StringMap field in the models.TemplateData struct and then it's used in the rendering of the html template for it's use in the webpage.

	remoteIP := repos.App.Session.GetString(r.Context(), "remote_ip")

	stringMap := make(map[string]string)
	stringMap["test"] = "hello again"
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
	message :=
		`
		The number of bytes written
		when accesing the about page
		is equal to %v
		`

	fmt.Print(message)
}
