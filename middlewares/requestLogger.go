package middlewares

import (
	"log"
	"net/http"
)

// global middleware for logging requests
func RequestLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s '%s' %s]", r.Method, r.URL, r.Proto)
		h.ServeHTTP(w, r)
	})
}