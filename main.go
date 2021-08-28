package main

import (
	"errors"
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

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 5.0
	y := 0.0
	z, err := divideValues(x, y)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", x, y, z)
}

func divideValues(x, y float64) (float64, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}

	return x / y, nil
}

func add(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}
