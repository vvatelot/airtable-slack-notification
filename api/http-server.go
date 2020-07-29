package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/vvatelot/airtable-slack-notify/thirdparty"
)

// RequestLogger is a logging middleware for http request
func RequestLogger(targetMux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		targetMux.ServeHTTP(w, r)

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

// CheckAirtableAllNewsHandler handle the /checknews route
func CheckAirtableAllNewsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	apiKey := query.Get("api_key")

	if apiKey == "" || apiKey != os.Getenv("API_KEY") {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Unauthorized"))
		return
	}

	go thirdparty.CheckAirableAllNews()
	fmt.Fprintf(w, "")
	log.Print("Call " + r.RequestURI)
}
