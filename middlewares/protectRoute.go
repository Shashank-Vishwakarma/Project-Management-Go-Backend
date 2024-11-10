package middlewares

import (
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/lib"
)

func VerifyToken(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			lib.HandleResponse(w, http.StatusInternalServerError, "Error reading cookie", nil)
			return
		}
		if token.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			lib.HandleResponse(w, http.StatusUnauthorized, "Unauthorized - Token not found", nil)
			return
		}

		err = lib.VerifyJWT(token.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			lib.HandleResponse(w, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func VerifyRole(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO
		h.ServeHTTP(w, r)
	})
}
