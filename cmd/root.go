package cmd

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/vvatelot/airtable-slack-notify/api"
	"github.com/vvatelot/airtable-slack-notify/thirdparty"
)

// Execute runs the root command
func Execute() {
	godotenv.Load()
	thirdparty.InitAirtableHTTPClient()

	thirdparty.Bases = strings.Split(os.Getenv("AIRTABLE_TABLES"), ",")

	mux := http.NewServeMux()

	mux.HandleFunc("/checknews", api.CheckAirtableAllNewsHandler)
	log.Fatal(http.ListenAndServe(":6060", api.RequestLogger(mux)))
}
