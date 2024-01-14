package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anjasfedo/go-postgres/routers"
)

// main function, the entry point of the program
func main() {
	// Obtain the router instance by calling the Router function from the routers package
	r := routers.Router()

	// Print a message indicating that the server is starting on port 8000
	fmt.Println("Start Server on Port 8000")

	// Start the HTTP server and listen on port 8000 using the obtained router
	log.Fatal(http.ListenAndServe(":8000", r))
}
