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

	if register.Name == "" || register.Email == "" || register.Role == "" {
		helpers.Response(w, 400, "Name, email, and role are required", nil)
		return
	}

	var existing models.User
	if err := configs.DB.Where("email = ?", register.Email).First(&existing).Error; err == nil {
		helpers.Response(w, 400, "Email is already registered", nil)
		return
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
		Role:     register.Role,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		log.Println("Database error:", err)
		helpers.Response(w, 500, "Internal server error", nil)
		return
	}

	helpers.Response(w, 201, "Registered successfully", nil)
}
