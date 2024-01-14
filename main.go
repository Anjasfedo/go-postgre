package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Anjasfedo/go-postgres/routers"
)

func main() {
	r := routers.Router()

	fmt.Println("Start Server on Port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}
