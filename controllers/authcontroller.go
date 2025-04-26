package controllers

import (
	"encoding/json"
	"go-auth-jwt/configs"
	"go-auth-jwt/helpers"
	"go-auth-jwt/models"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		log.Println("Decode error:", err)
		helpers.Response(w, 400, "Invalid input", nil)
		return
	}
	defer r.Body.Close()

	if register.Password != register.PasswordConfirm {
		helpers.Response(w, 400, "Password and Confirm Password do not match", nil)
		return
	}

	if register.Name == "" || register.Email == "" {
		helpers.Response(w, 400, "Name and Email are required", nil)
		return
	}

	var existing models.User
	if err := configs.DB.Where("email = ?", register.Email).First(&existing).Error; err == nil {
		helpers.Response(w, 400, "Email is already registered", nil)
		return
	}

	var userCount int64
	if err := configs.DB.Model(&models.User{}).Count(&userCount).Error; err != nil {
		log.Println("Count error:", err)
		helpers.Response(w, 500, "Internal server error", nil)
		return
	}

	var role string
	if userCount < 2 {
		role = "admin"
	} else {
		role = "user"
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		log.Println("Hash error:", err)
		helpers.Response(w, 500, "Internal server error", nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
		Role:     role,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		log.Println("Database error:", err)
		helpers.Response(w, 500, "Internal server error", nil)
		return
	}

	helpers.Response(w, 201, "Registered successfully", user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Login successfully", token)

}
