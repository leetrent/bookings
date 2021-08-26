package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the HOME page!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the ABOUT page (%d)!", add(2, 3))
}

func add(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}
