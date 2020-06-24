package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type slackRequestBody struct {
	Text string `json:"text"`
}

func generateMessage(items []item, base string) string {
	count := strconv.Itoa(len(items))
	message := "*" + count + " new " + base + "* submitted to Airtable. Please have a look! 🤪\n> \n"

	for _, item := range items {
		message += "> - *" + item.Fields.Name + "* (proposed by " + item.Fields.Contact_Mail + ")\n"
	}

	message += "\n> _Click here:_ https://airtable.com/" + os.Getenv("AIRTABLE_BASE")

	return message
}

func sendMessageToSlack(message string) error {
	slackBody, _ := json.Marshal(slackRequestBody{Text: message})
	req, err := http.NewRequest(http.MethodPost, os.Getenv("SLACK_WEBHOOK_URL"), bytes.NewBuffer(slackBody))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		log.Fatal(errors.New("Non-ok response returned from Slack"))
	}
	return nil
}
