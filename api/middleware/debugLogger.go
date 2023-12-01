package middleware

import (
	"fmt"
	"net/http"
)

func DebugLoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		fmt.Println("Request accepted:")

		fmt.Println("Route:", r.URL.Path)
		fmt.Println("Method:", r.Method)
		fmt.Println("Address:", r.RemoteAddr)

		fmt.Println("========================================================")

		next.ServeHTTP(w, r)
	})
}
