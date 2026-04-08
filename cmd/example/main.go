package main

import (
	"github.com/prashunchitkr/nepse-go/pkg/nepse"
)

func main() {
	nepseClient := nepse.NewClient()
	nepseClient.GetMarketSumary()
}
