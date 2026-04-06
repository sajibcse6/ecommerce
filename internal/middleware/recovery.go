package middleware

import (
	"log"
	"net/http"
	"runtime/debug"

	"ecommerce/internal/pkg/response"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				log.Printf("PANIC: %v\n%s", err, debug.Stack())

				response.Error(w, http.StatusInternalServerError, "Internal server error")
			}
		}()

		next.ServeHTTP(w, r)
	})
}