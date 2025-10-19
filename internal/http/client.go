// Package http
package http

import (
	"net/http"

	"github.com/go-resty/resty/v2"
	"github.com/prashunchitkr/nepse-go/internal/auth"
)

type HTTPClient struct {
	client *resty.Client
	auth   *auth.AuthHandler
}

func NewHTTPClient(inner *resty.Client, auth *auth.AuthHandler) *HTTPClient {
	c := &HTTPClient{
		client: inner,
		auth:   auth,
	}

	inner.OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		if req.URL == "/authenticate/prove" {
			return nil
		}

		token, err := c.auth.GetToken(req.Context())
		if err != nil {
			return err
		}

		req.SetHeader("Authorization", "Salter "+token.AccessToken)
		return nil
	})

	inner.OnAfterResponse(func(client *resty.Client, resp *resty.Response) error {
		if resp.StatusCode() == http.StatusUnauthorized {
			// TODO: Handle refresh logic
			return nil
		}

		// TODO: Handle retry logic

		return nil
	})

	return c
}

func (c *HTTPClient) Client() *resty.Client {
	return c.client
}
