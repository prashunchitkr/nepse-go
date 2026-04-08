package endpoints

import (
	"context"

	"github.com/prashunchitkr/nepse-go/internal/client"
	"github.com/prashunchitkr/nepse-go/internal/models"
)

type MarketService struct {
	client *client.Client
}

func (m *MarketService) GetSummary(ctx context.Context) ([]*models.Index, error) {
	var result []*models.Index

	err := m.client.Get(ctx, "/api/nots/nepse-index", &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
