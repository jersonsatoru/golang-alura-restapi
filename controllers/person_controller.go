package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jersonsatoru/golang-alura-restapi/database"
	"github.com/jersonsatoru/golang-alura-restapi/domain"
)

func GetPersons(w http.ResponseWriter, r *http.Request) {
	var p []domain.Person
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	qsPersonID := mux.Vars(r)["id"]
	personID, err := strconv.Atoi(qsPersonID)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}
	var person domain.Person
	database.DB.First(&person, personID)
	json.NewEncoder(w).Encode(person)
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person domain.Person
	json.NewDecoder(r.Body).Decode(&person)
	tx := database.DB.Create(&person)
	fmt.Println(tx.Error)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func DeletePersonByID(w http.ResponseWriter, r *http.Request) {
	personID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var person domain.Person
	tx := database.DB.First(&person, personID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	database.DB.Delete(&person, personID)
	w.WriteHeader(http.StatusNoContent)
}

func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	personID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var person domain.Person
	tx := database.DB.First(&person, personID)
	if tx.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(person)
	json.NewEncoder(w).Encode(person)
}
