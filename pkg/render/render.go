package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/3WDeveloper-GM/go-course-webapp/pkg/config"
	"github.com/3WDeveloper-GM/go-course-webapp/pkg/models"
)

// This variable is to make some functions accesible to the parse command, this functions can be used for formatting dates, making some unicode hacks, and other miscellaneous orders to be performed while rendering the html files.
var functions = template.FuncMap{}

// This variable acceses the AppConfig structure, the use of this variable is to have an app-wide confirguration across all the .go files, things like the TemplateCache, UseCache and other options are accessed with the help of this variable
var app *config.AppConfig

// This just retrieves the AppConfig variable from the main.go file and makes it so it's accesible in this file.
func NewTemplates(a *config.AppConfig) {
	app = a
}

// This is a function that will be used in the future, it just adds Default data to the .models.TemplateData instance in order to make it a more personalized experience for the user (this is the use that I could give to this function as of now 9/10/2023).
func AddDefaultData(TemplateData *models.TemplateData) *models.TemplateData {
	return TemplateData
}

// This is the most important function of the render.go package, this function parses the layout and template files that I created in the templates folder, and then renders them in the html in order to display the information that is located in the html file in the browser window.
func RenderTemplate(w http.ResponseWriter, tmpl string, TemplateData *models.TemplateData) {
	// I want to get the template cache from the configuration file, so i perform some operations in order to discern if I'm using the template cache or not.
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache // If I'm using the template cache, I just get the previously generated template cache from the main.go file and set it to the tc variable.
	} else {
		tc, _ = CreateTemplateCache() // If I'm not using the template cache (developer mode) this is the stuff I'll be using, the program generates the template cache every time the user acceses the website, this is slow and ineficcient, but can be useful for web development purposes.
	}

	//fmt.Println("passedcreatetemplatecache")

	// This logic block gets the template file that is used for the "tmpl" webpage. If there is not a template file that can be found in the template cache, the app is stopped using the log.Fatal command.
	t, ok := tc[tmpl]
	//fmt.Println(t)
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	//This buffer is created for adding data to the Web Response Writer. In this case, I think that this is meant to add the information of the buffer to the template.
	buf := new(bytes.Buffer)

	TemplateData = AddDefaultData(TemplateData)

	_ = t.Execute(buf, TemplateData)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

// Creates a template cache as a fast asf map
func CreateTemplateCache() (map[string]*template.Template, error) {

	// This is the variable that stores the templates in a cache.
	myCache := map[string]*template.Template{}
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//	fmt.Println("passed filepath.glob")

	// We range all the pages in the template directory. and then we match them with the templates.
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		//		fmt.Println("Passed filepath.base")

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		//		fmt.Println("passed matches")
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
			//			fmt.Println("passed parseglob")
		}
		myCache[name] = ts
	}
	return myCache, nil
}
