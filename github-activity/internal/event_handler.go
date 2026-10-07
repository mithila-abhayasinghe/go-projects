package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func FetchPublicRecordsByUser(uname string) {

	// api endpoint
	// https://api.github.com/users/kamranahmedse/events

	prod := true

	var events []GitHubEvent

	if prod {
		resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", uname))
		if err != nil {
			log.Fatal(err)
		}
		events, _ = parseEvents(resp.Body)
	}

	if !prod {
		file, _ := os.Open("my_events.json")
		defer file.Close()
		events, _ = parseEvents(file)
	}
	for i := range events {
		fmt.Println(events[i])
	}

}

func parseEvents(r io.Reader) ([]GitHubEvent, error) {
	var events []GitHubEvent
	err := json.NewDecoder(r).Decode(&events)
	return events, err
}
