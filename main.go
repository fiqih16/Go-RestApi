package main

import (
	"fmt"
	// "log"
	"net/http"

	admincontroller "github.com/fiqih16/Go-RestApi/controllers/admincontroller"
	"github.com/fiqih16/Go-RestApi/models"
	// "github.com/fiqih16/Go-RestApi/controllers/authcontroller"
	// "github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()
	// r := mux.NewRouter()

	// r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	// r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	// r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")
	
	// log.Fatal(http.ListenAndServe(":8080", r))

	http.HandleFunc("/", admincontroller.Index)
	http.HandleFunc("/dashboard", admincontroller.DAadmin)
	http.HandleFunc("/login", admincontroller.Login)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port 8080")
} // 34:01