package api

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "https://www.google.com")
		w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.Write([]byte("allowed"))
			return

		}

		next.ServeHTTP(w, r)
	})
}
