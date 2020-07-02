package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var httpClient http.Client

type item struct {
	AirtableID string `json:"id"`
	Fields     struct {
		Name        string `json:"name"`
		ContactMail string `json:"contact_mail"`
	} `json:"fields"`
}

type responseBody struct {
	Records []item `json:"records"`
}

func initHTTPClient() {
	tr := &http.Transport{
		TLSHandshakeTimeout: 30 * time.Second,
	}
	httpClient = http.Client{Transport: tr}
}

func genereteRowQuery(req *http.Request) string {
	q := req.URL.Query()
	q.Add("fields[]", "name")
	q.Add("fields[]", "contact_mail")
	q.Add("filterByFormula", "{status}='proposal'")
	return q.Encode()
}

func getNewItemsHTTPClient(table string) []item {
	path := []string{os.Getenv("AIRTABLE_API_URL"), os.Getenv("AIRTABLE_BASE"), table}
	url := strings.Join(path, "/")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer "+os.Getenv("AIRTABLE_API_KEY"))

	req.URL.RawQuery = genereteRowQuery(req)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err)
	}

	var r responseBody
	jsonError := json.Unmarshal(body, &r)
	if jsonError != nil {
		log.Panic(jsonError)
	}
	return r.Records
}

func checkAllNews() {
	for _, base := range bases {
		newItems := getNewItemsHTTPClient(base)
		if len(newItems) > 0 {
			message := generateMessage(newItems, base)
			log.Println(message)
			sendMessageToSlack(message)
		} else {
			log.Println("No new " + base + " items")
		}
	}
}
