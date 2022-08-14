package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/fiqih16/Go-RestApi/config"
	"github.com/fiqih16/Go-RestApi/helper"
	"github.com/fiqih16/Go-RestApi/models"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// menggambil input json dari request
	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response) 
		return
	}
	defer r.Body.Close()

	// ambil data user berdasarkan email
	var user models.User
	if err := models.DB.Where("email = ?", userInput.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "Email atau password salah"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}

	// cek apakah password valid
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		response := map[string]string{"message": "Email atau password salah"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	// jika berhasil login, generate token (buat token jwt)
	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTclaim{
		Username : user.Username,
		RegisteredClaims : jwt.RegisteredClaims{
			Issuer: "Go-RestApi",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// generate token / signed token
	token, err := tokenAlgo.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set token ke cookie
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Path: "/",
		Value: token,
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login Berhasil"}
	helper.ResponseJSON(w, http.StatusOK, response)
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
	}
	

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