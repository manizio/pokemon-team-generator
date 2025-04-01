package middleware

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("\tRequest: %s %s\n\n", r.Method, r.URL.Path)
		fmt.Println("\tForm data: ", r.Form, "\n\n")
		next.ServeHTTP(w, r)
	}) 
	
}
