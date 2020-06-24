package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func requestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

		// log request by who(IP address)
		requesterIP := r.RemoteAddr

		log.Printf(
			"%s\t\t%s\t\t%s\t\t%v",
			r.Method,
			r.RequestURI,
			requesterIP,
			time.Since(start),
		)
	})
}

func checkAllNewsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	apiKey := query.Get("api_key")

	if apiKey == "" || apiKey != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unauthorized"))
		return
	}

	go checkAllNews()
	fmt.Fprintf(w, "")
	log.Print("Call " + r.RequestURI)
}
