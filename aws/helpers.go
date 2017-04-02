package main

import (
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
  "log"
)

// Wrapper for normal logging requests
// For internal endpoints, only outputs to log file
func Logger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(">>>", r.Method, r.RequestURI)
		inner.ServeHTTP(w, r)
	})
}

// Wrapper for CORS OPTIONS request
// Writes a CORS header to requests and handle the OPTIONS method
func HandleOptions(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			writeCORSHeader(w, r)
			return
		} else {
			writeCORSHeader(w, r)
			h.ServeHTTP(w, r)
		}
	}
}

// Writes the CORS response header
func writeCORSHeader(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	if origin == "" {
		origin = "*"
	}
	w.Header().Set("Access-Control-Allow-Origin", origin)      // SSL requires matching origin, * does not work with SSL
	w.Header().Set("Access-Control-Allow-Credentials", "true") // enable SSL
	w.Header().Set("Access-Control-Expose-Headers", "Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PATCH, PUT, DELETE")
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}

// Gets an integer parameter from the URL
func getIntParam(r *http.Request, paramName string) int {
	vars := mux.Vars(r)
	if v, ok := vars[paramName]; ok {
		if i, err := strconv.Atoi(v); err != nil {
			log.Printf("param %v with value %v is not an integer", paramName, v)
			return 1
		} else {
			return i
		}
	} else {
		log.Printf("param %v does not exist\n", paramName)
		return 1
	}
}

// Gets a string parameter from the URL
func getStrParam(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	if v, ok := vars[paramName]; ok {
		return v
	} else {
		log.Printf("param %v does not exist\n", paramName)
		return ""
	}
}
