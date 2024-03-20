package middleware

import (
	"net/http"
)

func AllowOnlyFrom(allowedDomain string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin != allowedDomain {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		handler.ServeHTTP(w, r)
	})
}
