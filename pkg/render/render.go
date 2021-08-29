package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(rw http.ResponseWriter, fileName string) {
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
