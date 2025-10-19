package apitypes

const (
	ACTIVE   = 'A'
	INACTIVE = "I"
)

type Security struct {
	ID           uint64 `json:"id"`
	Symbol       string `json:"symbol"`
	SecurityName string `json:"securityName"`
	Name         string `json:"name"`
	ActiveStatus string `json:"activeStatus"`
}
