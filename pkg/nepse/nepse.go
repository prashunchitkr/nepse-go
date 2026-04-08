// Package nepse
package nepse

type Option func(client *Client)

type Client struct{}

func NewClient(options ...Option) (*Client, error) {
	return &Client{}, nil
}
