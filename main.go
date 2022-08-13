package main

import (
	"log"
	"net/http"

	"github.com/fiqih16/Go-RestApi/controllers/authcontroller"
	"github.com/fiqih16/Go-RestApi/models"

	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", r))
}