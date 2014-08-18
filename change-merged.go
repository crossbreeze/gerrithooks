package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// change this setting
const slackWebHookURL = "Slack Webhook URL"

type payload struct {
	Text string `json:"text"`
}

// change-merged --change <change id> --change-url <change url> --project <project name> --branch <branch> --topic <topic> --submitter <submitter> --commit <sha1>
func main() {
	var change string
	var changeURL string
	var project string
	var branch string
	var topic string
	var submitter string
	var commit string

	flag.StringVar(&change, "change", "", "change id")
	flag.StringVar(&changeURL, "change-url", "", "change url")
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&branch, "branch", "", "branch")
	flag.StringVar(&topic, "topic", "", "topic")
	flag.StringVar(&submitter, "submitter", "", "submitter")
	flag.StringVar(&commit, "commit", "", "sha1")

	flag.Parse()

	p := payload{fmt.Sprintf("%s has merged a commit (%s) at <%s> for %s (branch: %s)", submitter, commit, changeURL, project, branch)}

	payloadJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Post(slackWebHookURL, "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		log.Fatalln(err)
	}
}
