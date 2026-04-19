// Package client
package client

import (
	"context"
	"encoding/json"
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

func NewClient(httpClient *http.Client, session *session.Manager, auth *auth.Manager, baseURL string) *Client {
	return &Client{
		httpClient: httpClient,
		session:    session,
		auth:       auth,
		baseURL:    baseURL,
	}
}

func (c *Client) Get(ctx context.Context, endpoint string, result any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.baseURL+endpoint, http.NoBody)
	if err != nil {
		return err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status code: " + resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}
