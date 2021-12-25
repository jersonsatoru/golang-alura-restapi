package main

import (
	"log"
	"net/http"

	"github.com/jersonsatoru/golang-alura-restapi/database"
	"github.com/jersonsatoru/golang-alura-restapi/routes"
)

func main() {
	r := routes.HandleRoutes()
	database.Connect()
	log.Fatal(http.ListenAndServe(":8080", r))
}
