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

	summary, err := client.GetSummary(context.Background())
	if err != nil {
		log.Fatalln("error getting summary:", err)
	}

	for _, index := range summary {
		log.Printf("%+v", index)
	}
}
