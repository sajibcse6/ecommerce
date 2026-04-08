package middleware

import (
	"net/http"
	"strings"

	"ecommerce/internal/pkg/response"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(secret string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")

			if authHeader == "" {
				response.Error(w, http.StatusUnauthorized, "Missing token")
				return
			}

			tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

			token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(secret), nil
			})

			if err != nil || !token.Valid {
				response.Error(w, http.StatusUnauthorized, "Invalid token")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}