// Package client
package client

import (
	"context"
	"errors"
	"net/http"

	"github.com/prashunchitkr/nepse-go/internal/auth"
	"github.com/prashunchitkr/nepse-go/internal/session"
)

type Client struct {
	httpClient *http.Client
	session    *session.Manager
	auth       *auth.Manager
	baseURL    string
}

func (c *Client) Get(ctx context.Context, endpoint string, result any) error {
	return errors.New("not implemented")
}
