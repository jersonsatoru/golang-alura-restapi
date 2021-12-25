package routes

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/jersonsatoru/golang-alura-restapi/controllers"
	"github.com/jersonsatoru/golang-alura-restapi/middlewares"
)

func HandleRoutes() http.Handler {
	r := mux.NewRouter()
	r.Use(middlewares.JsonContentType)
	r.HandleFunc("/persons", controllers.GetPersons).Methods(http.MethodGet)
	r.HandleFunc("/personsc", controllers.CreatePerson).Methods(http.MethodPost)
	r.HandleFunc("/persons/{id}", controllers.GetPersonByID).Methods(http.MethodGet)
	r.HandleFunc("/persons/{id}", controllers.DeletePersonByID).Methods(http.MethodDelete)
	r.HandleFunc("/persons/{id}", controllers.UpdatePersonByID).Methods(http.MethodPut)
	return handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)
}
