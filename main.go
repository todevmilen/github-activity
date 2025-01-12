package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Event struct {
	Id   json.Number `json:"id"`
	Type string      `json:"type"`
	Repo struct {
		Id   json.Number `json:"id"`
		Name string      `json:"name"`
	} `json:"repo"`
	Payload struct {
		Size json.Number `json:"size"`
	} `json:"payload"`
}

func main() {
	var events []Event

	if len(os.Args) <= 1 {
		log.Fatal("Please provide username")
	}

	username := os.Args[1]

	if username == "" {
		log.Fatal("Please provide username")
	}

	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode == 404 {
		log.Fatalln("Could not find repository")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &events)
	if err != nil {
		log.Fatalln(err)
	}

	for _, event := range events {
		switch event.Type {
		case "PushEvent":
			fmt.Printf("- Pushed %s commits to %s\n", event.Payload.Size, event.Repo.Name)

		case "ForkEvent":
			fmt.Printf("- Forked  %s\n", event.Repo.Name)

		case "WatchEvent":
			fmt.Printf("- Starred the repository %s\n", event.Repo.Name)

		default:
			fmt.Printf("- Some Event into %s\n", event.Repo.Name)
		}

	}

}
