// Package nepsego is a wrapper for NEPSE REST API
package nepsego

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prashunchitkr/nepse-go/internal/auth"
	internalHttp "github.com/prashunchitkr/nepse-go/internal/http"
)

const (
	baseURL = "https://nepalstock.com.np/api"
)

type Client struct {
	httpClient *internalHttp.HTTPClient
	wasmHelper *auth.WasmHelper
}

func NewClient() *Client {
	ctx := context.Background()
	restyClient := resty.New().
		SetTLSClientConfig(&tls.Config{
			InsecureSkipVerify: true,
		}).
		SetTimeout(5 * time.Second).
		SetBaseURL(baseURL).
		SetHeaders(map[string]string{
			"User-Agent":   "Mozilla/5.0 (X11; Linux i686; rv:136.0) Gecko/20100101 Firefox/136.0",
			"Content-Type": "application/json",
			"Accept":       "application/json",
		})

	wasmHelper, err := auth.NewWasmHelper(ctx)
	if err != nil {
		log.Fatalf("Error initializing wasm: %v", err)
	}

	authHandler := auth.NewAuthHandler(*restyClient, wasmHelper)

	internalHTTPClient := internalHttp.NewHTTPClient(restyClient, authHandler)

	return &Client{
		httpClient: internalHTTPClient,
		wasmHelper: wasmHelper,
	}
}

func (c *Client) Close() {
	c.wasmHelper.Close()
}
