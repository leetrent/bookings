package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const port = ":8080"

func Home(rw http.ResponseWriter, req *http.Request) {
	renderTemplate(rw, "home.page.tmpl")
}

func About(rw http.ResponseWriter, req *http.Request) {
	renderTemplate(rw, "about.page.tmpl")
}

func renderTemplate(rw http.ResponseWriter, fileName string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + fileName)
	if err != nil {
		fmt.Printf("error encountered when invoking template.ParseFiles for %s", fileName)
		return
	}

	err2 := parsedTemplate.Execute(rw, nil)
	if err2 != nil {
		fmt.Printf("error encountered when invoking parsedTemplate.Execute for %s", fileName)
		return
	}

}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}
