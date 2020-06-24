package main

import (
	"log"
	"os"

	"github.com/fabioberger/airtable-go"
)

var airtableClient *airtable.Client
var airtableAPIKey, baseID string

type item struct {
	AirtableID string
	Fields     struct {
		Name         string
		Contact_Mail string
	}
}

func initAirtableClient() {
	var err error
	airtableClient, err = airtable.New(os.Getenv("AIRTABLE_API_KEY"), os.Getenv("AIRTABLE_BASE"))
	if err != nil {
		log.Panic(err)
	}
}

func getNewItems(base string) []item {
	listParams := airtable.ListParameters{
		Fields:          []string{"name", "contact_mail"},
		FilterByFormula: "{status}='proposal'",
	}

	items := []item{}
	if err := airtableClient.ListRecords(base, &items, listParams); err != nil {
		log.Panic(err)
	}

	return items
}
