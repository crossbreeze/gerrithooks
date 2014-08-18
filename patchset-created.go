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

// patchset-created --change <change id> --is-draft <boolean> --change-url <change url> --project <project name> --branch <branch> --topic <topic> --uploader <uploader> --commit <sha1> --patchset <patchset id>
func main() {
	var change string
	var draft string
	var changeURL string
	var project string
	var branch string
	var topic string
	var uploader string
	var commit string
	var patchset string

	flag.StringVar(&change, "change", "", "change id")
	flag.StringVar(&draft, "is-draft", "", "draft")
	flag.StringVar(&changeURL, "change-url", "", "change url")
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&branch, "branch", "", "branch")
	flag.StringVar(&topic, "topic", "", "topic")
	flag.StringVar(&uploader, "uploader", "", "uploader")
	flag.StringVar(&commit, "commit", "", "sha1")
	flag.StringVar(&patchset, "patchset", "", "patchset id")

	flag.Parse()

	p := payload{fmt.Sprintf("%s has created a new patchset (%s) at <%s> for %s (branch: %s)", uploader, patchset, changeURL, project, branch)}

	payloadJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Post(slackWebHookURL, "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		log.Fatalln(err)
	}
}
