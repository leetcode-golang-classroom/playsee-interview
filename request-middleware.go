package main

import (
	"net/http"
)

var (
	errorApiKeyNotProvide = "header api-key not provide"
	errorForWrongApiKey   = "wrong api key"
)

func RequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// parse request header has request key 'api-key'
		apiKey := r.Header.Get("api-key")
		if apiKey == "" {
			http.Error(w, errorApiKeyNotProvide, http.StatusUnauthorized)
			return
		}
		if apiKey != APIKey {
			http.Error(w, errorForWrongApiKey, http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
