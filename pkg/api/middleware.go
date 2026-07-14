package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		password := os.Getenv("TODO_PASSWORD")

		if password == "" {
			next(w, r)
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			http.Error(w, "Авторизация провалена", http.StatusUnauthorized)
			return
		}

		if !validateToken(cookie.Value, password) {
			http.Error(w, "Авторизация провалена", http.StatusUnauthorized)
			return
		}

		next(w, r)
	}
}

func validateToken(tokenString, password string) bool {

	token, err := jwt.Parse(
		tokenString,
		func(token *jwt.Token) (any, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return []byte(jwtSecret), nil
		},
	)

	if err != nil {
		return false
	}

	if !token.Valid {
		return false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}

	hash, ok := claims["hash"].(string)
	if !ok {
		return false
	}

	return hash == createHash(password)
}
