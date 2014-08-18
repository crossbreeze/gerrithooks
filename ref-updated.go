package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	// change these settings
	slackWebHookURL = "Slack Webhook URL"
	diffURL         = "http://www.gerrit.net/gitweb?p=%s.git;a=commitdiff;h=%s"
)

type payload struct {
	Text string `json:"text"`
}

// ref-updated --oldrev <old rev> --newrev <new rev> --refname <ref name> --project <project name> --submitter <submitter>
func main() {
	var oldrev string
	var newrev string
	var refname string
	var project string
	var submitter string

	flag.StringVar(&oldrev, "oldrev", "", "old rev")
	flag.StringVar(&newrev, "newrev", "", "new rev")
	flag.StringVar(&refname, "refname", "", "ref name")
	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&submitter, "submitter", "", "submitter")

	flag.Parse()

	diffURL := fmt.Sprintf(diffURL, project, newrev)
	text := fmt.Sprintf("%s has updated %s (%s) to %s - <%s|diff>", submitter, project, refname, newrev, diffURL)

	p := payload{text}

	payloadJSON, err := json.Marshal(p)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = http.Post(slackWebHookURL, "application/json", bytes.NewReader(payloadJSON))
	if err != nil {
		log.Fatalln(err)
	}
}
