package nepsego

import (
	"context"
	"errors"

	apitypes "github.com/prashunchitkr/nepse-go/internal/types"
)

func (c *Client) GetSecurities(ctx context.Context) (*[]apitypes.Security, error) {
	var securities []apitypes.Security

	resp, err := c.httpClient.Client().R().
		SetContext(ctx).
		SetResult(&securities).
		SetQueryParam("nonDelisted", "true").
		Get("/nots/security")
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, errors.New("getting securities failed: " + resp.Status())
	}

	return &securities, nil
}
