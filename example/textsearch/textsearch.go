package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"gopkg.in/maxhawkins/google-places-api.v2/places"
)

func main() {
	key := flag.String("key", "", "google places api key")
	flag.Parse()

	service := places.NewService(http.DefaultClient, *key)

	call := service.TextSearch("Google")

	resp, err := call.Do()
	if places.IsZeroResults(err) {
		fmt.Println("no results")
		return
	}
	if err != nil {
		panic(err)
	}

	results := resp.Results
	token := resp.NextPageToken

	for token != "" {
		time.Sleep(2 * time.Second) // Rate limit

		call.PageToken = token
		resp, err := call.Do()
		if err != nil {
			panic(err)
		}

		token = resp.NextPageToken
		results = append(results, resp.Results...)
	}

	for _, result := range results {
		fmt.Println(result.Name)
	}
}
