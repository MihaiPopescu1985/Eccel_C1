package controller

import (
	"log"
	"net/http"
)

func LogInMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "no-cache, no-store, must-revalidate")
		log.Println(r.RequestURI)

		if r.RequestURI != "/" {
			http.Redirect(w, r, "/", 300)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
