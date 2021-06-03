package middleware

import (
	"log"
	"net/http"
	"os"
)

func ResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adds content-type of application/json to Header
		w.Header().Add("Content-Type", "application/json")
		// These headers help prevent XSS and Clickjacking attacks
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("X-Frame-Options", "deny")
		next.ServeHTTP(w, r)
	})
}

func LogRequest(next http.Handler) http.Handler {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		infoLog.Printf("%s %s %s", r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
