package admincontroller

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/home/index.html")
	temp.Execute(w, nil)
}

func DAadmin(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/dashboard/dashboard.html")
	temp.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	temp, _ := template.ParseFiles("views/login/login.html")
	temp.Execute(w, nil)
}