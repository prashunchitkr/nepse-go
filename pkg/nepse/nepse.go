// Package nepse
package nepse

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/prashunchitkr/nepse-go/internal/auth"
	"github.com/prashunchitkr/nepse-go/internal/client"
	"github.com/prashunchitkr/nepse-go/internal/endpoints"
	"github.com/prashunchitkr/nepse-go/internal/session"
)

type Option func(client *Client)

type Client struct {
	marketService *endpoints.MarketService
}

func NewClient(options ...Option) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{
		Timeout:   5 * time.Second,
		Transport: tr,
	}
	session := session.NewManager()
	auth := auth.NewManager(session)
	c := client.NewClient(httpClient, session, auth, "https://nepalstock.com.np")

	markerService := endpoints.NewMarketService(c)

	return &Client{
		marketService: markerService,
	}, nil
}
