package main

import (
	"log"

	"github.com/prashunchitkr/nepse-go/pkg/nepse"
)

func main() {
	_, err := nepse.NewClient()
	if err != nil {
		log.Fatalln("error initializing client")
	}

	log.Println("Client Initialized")
}
