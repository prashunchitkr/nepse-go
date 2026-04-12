package models

type Security struct {
	ID           uint32 `json:"id"`
	Symbol       string `json:"symbol"`
	SecurityName string `json:"securityName"`
	Name         string `json:"name"`
	ActiveStatus string `json:"activeStatus"`
}
