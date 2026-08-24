package main

import "net/http"

func enableCORS(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Allow-Control-Origin", "*")
		w.Header().Set("Access-Allow-Control-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Allow-Control-Headers", "Content-Type, Authorization")

		// Allow preflight requests from the browser API
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler(w, r)
	}
}
