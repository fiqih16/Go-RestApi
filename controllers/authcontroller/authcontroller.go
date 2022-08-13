package authcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/fiqih16/Go-RestApi/helper"
	"github.com/fiqih16/Go-RestApi/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, r *http.Request) {

}

func Register(w http.ResponseWriter, r *http.Request) {
	// menggambil input json dari request
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response) 
		return
	}
	defer r.Body.Close()

	// jika email sudah terdaftar tidak bisa register
	if models.IsEmailExist(userInput.Email) {
		response := map[string]string{"message": "Email sudah terdaftar"}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	} // 20:24
	

	// hash password menggunakan bcrypt
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	// insert ke database
	if err := models.DB.Create(&userInput).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response) 
		return
	}

	response := map[string]string{"message": "Berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response) 


}

func Logout(w http.ResponseWriter, r *http.Request) {

}