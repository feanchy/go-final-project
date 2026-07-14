package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type signinRequest struct {
	Password string `json:"password"`
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("signin")

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req signinRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	originPass := os.Getenv("TODO_PASSWORD")
	fmt.Println(originPass)
	if req.Password != originPass {
		writeJSON(w, map[string]string{
			"error": "Неверный пароль",
		})
		return
	}

	token, err := createToken(originPass)
	if err != nil {
		writeJSON(w, map[string]string{
			"error": err.Error(),
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		MaxAge:   8 * 60 * 60,
		HttpOnly: true,
	})

	writeJSON(w, map[string]string{
		"token": token,
	})
}
