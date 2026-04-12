package endpoints

import (
	"context"
	"net/url"
	"strconv"

	"github.com/prashunchitkr/nepse-go/internal/client"
	"github.com/prashunchitkr/nepse-go/internal/models"
)

type MarketService struct {
	client *client.Client
}

func NewMarketService(c *client.Client) *MarketService {
	return &MarketService{
		client: c,
	}
}

func (m *MarketService) GetNepseIndex(ctx context.Context) ([]*models.NepseIndex, error) {
	var result []*models.NepseIndex

	err := m.client.Get(ctx, "/api/nots/nepse-index", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (m *MarketService) GetSecurities(ctx context.Context, nonDelisted bool) ([]*models.Security, error) {
	var result []*models.Security

	query := url.Values{}
	query.Add("nonDelisted", strconv.FormatBool(nonDelisted))

	err := m.client.Get(ctx, "/api/nots/security?"+query.Encode(), &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
