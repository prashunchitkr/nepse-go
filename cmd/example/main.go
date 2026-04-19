package main

import (
	"context"
	"log"

	"github.com/prashunchitkr/nepse-go/pkg/nepse"
)

func main() {
	client, err := nepse.NewClient()
	if err != nil {
		log.Fatalln("error initializing client")
	}

	log.Println("Client Initialized")

	summary, err := client.GetNepseIndex(context.Background())
	if err != nil {
		log.Fatalln("error getting summary:", err)
	}

	for _, index := range summary {
		log.Printf("%s: %f", index.Index, index.CurrentValue)
	}

	securities, err := client.GetSecurities(context.Background(), false)
	if err != nil {
		log.Fatalln("error getting securities:", err)
	}

	for _, security := range securities {
		log.Printf("%s", security.Name)
	}
}
