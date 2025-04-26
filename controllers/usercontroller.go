package controllers

import (
	"go-auth-jwt/helpers"
	"go-auth-jwt/models"
	"net/http"
)

func Me(w http.ResponseWriter, r *http.Request) {
	userInfo := r.Context().Value("userinfo")
	claims, ok := userInfo.(*helpers.Claims)
	if !ok || claims == nil {
		helpers.Response(w, 401, "Unauthorized", nil)
		return
	}

	userResponse := &models.MyProfile{
		ID:    claims.ID,
		Name:  claims.Name,
		Email: claims.Email,
		Role:  claims.Role,
	}

	helpers.Response(w, 200, "My Profile", userResponse)
}
