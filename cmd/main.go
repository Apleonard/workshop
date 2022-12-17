package main

import (
	"net/http"
	"workshop/internal"
	"workshop/models"

	"github.com/gorilla/mux"
)

func main() {
	models.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/test", internal.Test).Methods("GET")

	// create person
	r.HandleFunc("/create", internal.Create).Methods("POST")

	// get list person
	// r.HandleFunc("/people", internal.GetPeople).Methods("GET")

	// get detail person
	r.HandleFunc("/person/{id}", internal.GetPerson).Methods("GET")

	// update person
	r.HandleFunc("/update-person/{id}", internal.UpdatePerson).Methods("PUT")

	// delete person
	r.HandleFunc("/delete-person/{id}", internal.DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8080", r)

}
