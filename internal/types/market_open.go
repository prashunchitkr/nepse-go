package apitypes

type MarketOpen struct {
	IsOpen string `json:"isOpen"`
	AsOf   string `json:"asOf"`
	ID     int    `json:"id"`
}
