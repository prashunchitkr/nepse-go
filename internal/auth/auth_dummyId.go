package auth

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	apitypes "github.com/prashunchitkr/nepse-go/internal/types"
)

var lut = []int{
	147, 117, 239, 143, 157, 312, 161, 612, 512, 804, 411, 527, 170, 511, 421,
	667, 764, 621, 301, 106, 133, 793, 411, 511, 312, 423, 344, 346, 653, 758,
	342, 222, 236, 811, 711, 611, 122, 447, 128, 199, 183, 135, 489, 703, 800,
	745, 152, 863, 134, 211, 142, 564, 375, 793, 212, 153, 138, 153, 648, 611,
	151, 649, 318, 143, 117, 756, 119, 141, 717, 113, 112, 146, 162, 660, 693,
	261, 362, 354, 251, 641, 157, 178, 631, 192, 734, 445, 192, 883, 187, 122,
	591, 731, 852, 384, 565, 596, 451, 772, 624, 691,
}

func extractDate(dateTimeStr string) (int, error) {
	layout := "2006-01-02T15:00:00"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return -1, err
	}

	return t.Day(), nil
}

func (a *AuthHandler) getDummyID(ctx context.Context, accessToken string) (int, error) {
	log.Printf("Getting dummy id")
	var marketOpen apitypes.MarketOpen

	resp, err := a.client.R().
		SetContext(ctx).
		SetResult(&marketOpen).
		SetHeader("Authorization", "Salter "+accessToken).
		Get("/nots/nepse-data/market-open")
	if err != nil {
		return -1, err
	}

	if resp.IsError() {
		return -1, errors.New("getting market status failed: " + resp.Status())
	}

	serverDay, err := extractDate(marketOpen.AsOf)
	if err != nil {
		return -1, fmt.Errorf("error parsing date: %v", err)
	}

	dummyID := lut[marketOpen.ID] + marketOpen.ID + 2*serverDay

	return dummyID, nil
}
