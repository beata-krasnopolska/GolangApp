package handlers

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

//in order function handle requests from browser it has to have 2 parameters: w http.ResponseWriter, r *http.Request

func Home(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.page.html")
}

// package lever variable for template cache
var tc = make(map[string]*template.Template)

func renderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.html")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we alredy have the template in our chache
	_, inMap := tc[t]
	if !inMap {
		// need to create the template
		log.Println("Creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the templete in the cache
		log.Println("using cached template")
	}

	tmpl = tc[t]

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.html",
	}

	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// add template to cache
	tc[t] = tmpl

	return nil
}
