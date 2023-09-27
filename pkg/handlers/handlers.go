package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

//in order function handle requests from browser it has to have 2 parameters: w http.ResponseWriter, r *http.Request

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.html")
}

// small letter in function name is like making it private
func addValues(x, y int) int {
	return x + y
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template: ", err)
		return
	}
}
