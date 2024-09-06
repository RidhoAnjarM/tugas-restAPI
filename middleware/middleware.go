package middleware

import (
	"encoding/json"
	"net/http"
	"strings"
)

func TokenValidation(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") || authHeader[7:] != "token_app" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"status":  false,
				"message": "Unauthorized",
				"data":    nil,
			})
			return
		}
		next.ServeHTTP(w, r)
	}
}
