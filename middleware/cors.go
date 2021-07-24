package middleware

import "net/http"

//CorHandler enables cors from all origins
func CorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		origin := r.Header.Get("Origin")

		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Origin", origin)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
		} else {
			next.ServeHTTP(w, r)
		}

	})
}
