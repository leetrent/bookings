package main

import (
	"fmt"
	"net/http"

	"github.com/leetrent/bookings/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s", port)
	http.ListenAndServe(port, nil)
}
