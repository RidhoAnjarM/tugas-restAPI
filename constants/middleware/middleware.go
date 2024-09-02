package middleware

import (
    "net/http"
)

// TokenValidation: Memvalidasi token dari header Authorization
func TokenValidation(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token != "Bearer token_app" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }
        next.ServeHTTP(w, r)
    }
}
