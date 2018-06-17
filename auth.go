package main

import (
	"log"
	"net/http"
)

func getAuth(r *http.Request) *string {
	token := r.Header.Get("Authorization")
	if len(token) < 1 {
		return nil
	}
	return &token
}

func use(h http.HandlerFunc, middleware ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for _, m := range middleware {
		h = m(h)
	}

	return h
}

func firebaseAuth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.Header().Set("Content-Type", "application/json")
		if getAuth(r) == nil {
			log.Println("Not authorized")
			http.Error(w, "Not authorized", 401)
			return
		}
		h.ServeHTTP(w, r)
	}
}
