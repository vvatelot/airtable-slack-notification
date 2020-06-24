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
	initAirtableClient()

	bases = strings.Split(os.Getenv("AIRTABLE_TABLES"), ",")

	mux := http.NewServeMux()

	mux.HandleFunc("/checknews", checkAllNewsHandler)
	log.Fatal(http.ListenAndServe(":8080", requestLogger(mux)))
}

func checkAllNews() {
	for _, base := range bases {
		newItems := getNewItems(base)
		if len(newItems) > 0 {
			message := generateMessage(newItems, base)
			log.Println(message)
			sendMessageToSlack(message)
		}
	}
}
