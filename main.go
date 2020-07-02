package main

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var bases []string

func main() {
	godotenv.Load()
	initHTTPClient()

	bases = strings.Split(os.Getenv("AIRTABLE_TABLES"), ",")

	mux := http.NewServeMux()

	mux.HandleFunc("/checknews", checkAllNewsHandler)
	log.Fatal(http.ListenAndServe(":6060", requestLogger(mux)))
}
