// Package nepse
package nepse

import (
	"context"
	"crypto/tls"
	"net/http"
	"time"

	"github.com/prashunchitkr/nepse-go/internal/auth"
	"github.com/prashunchitkr/nepse-go/internal/client"
	"github.com/prashunchitkr/nepse-go/internal/endpoints"
	"github.com/prashunchitkr/nepse-go/internal/models"
	"github.com/prashunchitkr/nepse-go/internal/session"
	"github.com/prashunchitkr/nepse-go/internal/transport"
)

const defaultBaseURL = "https://nepalstock.com.np"

type Option func(client *Client)

type Client struct {
	marketService *endpoints.MarketService
}

func NewClient(options ...Option) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	sess := session.NewManager()
	publicHTTP := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport.NewNepsePublicTransport(tr),
	}
	authMgr := auth.NewManager(sess, publicHTTP, defaultBaseURL)
	apiHTTP := &http.Client{
		Timeout:   5 * time.Second,
		Transport: transport.NewNepseAuthTransport(tr, authMgr),
	}
	c := client.NewClient(apiHTTP, sess, authMgr, defaultBaseURL)

	markerService := endpoints.NewMarketService(c)

	return &Client{
		marketService: markerService,
	}, nil
}

func (c *Client) GetSummary(ctx context.Context) ([]*models.Index, error) {
	return c.marketService.GetSummary(ctx)
}
