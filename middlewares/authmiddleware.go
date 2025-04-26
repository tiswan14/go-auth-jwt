package middlewares

import (
	"context"
	"go-auth-jwt/helpers"
	"net/http"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		acccessToken := r.Header.Get("Authorization")

		if acccessToken == "" {
			helpers.Response(w, 401, "Unauthorized", "Please login first")
			return
		}

		user, err := helpers.ValidateToken(acccessToken)

		if err != nil {
			helpers.Response(w, 401, err.Error(), "Please login first")
			return
		}
		ctx := context.WithValue(r.Context(), "userinfo", user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
