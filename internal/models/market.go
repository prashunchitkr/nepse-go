package models

// Exchange Index IDs
const (
	SensitiveIdxID      = 57
	NepseIdxID          = 58
	FloatIdxID          = 62
	FloatSensitiveIdxID = 63
)

type NepseIndex struct {
	ID               uint32  `json:"id"`
	AuditID          any     `json:"auditId"`
	ExchangeIndexID  any     `json:"exchangeIndexId"`
	GeneratedTime    string  `json:"generatedTime"`
	Index            string  `json:"index"`
	Close            float32 `json:"close"`
	High             float32 `json:"high"`
	Low              float32 `json:"low"`
	PreviousClose    float32 `json:"previousClose"`
	Change           float32 `json:"change"`
	PerChange        float32 `json:"perChange"`
	FiftyTwoWeekHigh float32 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow  float32 `json:"fiftyTwoWeekLow"`
	CurrentValue     float32 `json:"currentValue"`
}
