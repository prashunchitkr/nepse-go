package main

import (
	"context"
	"fmt"
	"log"

	"github.com/prashunchitkr/nepse-go/pkg/nepsego"
)

func main() {
	fmt.Println("Nepse Go example")
	nepse := nepsego.NewClient()

	defer nepse.Close()

	ctx := context.Background()
	securities, err := nepse.GetSecurities(ctx)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	log.Printf("Retrieved %d securities", len(*securities))
}
