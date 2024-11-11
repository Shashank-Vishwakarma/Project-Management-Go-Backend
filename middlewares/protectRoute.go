package middlewares

import (
	"context"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/constants"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/lib"
)

func VerifyToken(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			lib.HandleResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		if token.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			lib.HandleResponse(w, http.StatusUnauthorized, "Unauthorized - Token not found", nil)
			return
		}

		claims, err := lib.VerifyJWT(token.Value)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			lib.HandleResponse(w, http.StatusUnauthorized, err.Error(), nil)
			return
		}

		ctx := context.WithValue(r.Context(), constants.USER_CONTEXT_KEY, claims)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ChainMiddlewares(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}