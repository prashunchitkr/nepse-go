// Package apitypes
package apitypes

type Prove struct {
	ServerTime      uint64 `json:"serverTime"`
	Salt            string `json:"salt"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
	TokenType       string `json:"tokenType"`
	Salt1           uint64 `json:"salt1"`
	Salt2           uint64 `json:"salt2"`
	Salt3           uint64 `json:"salt3"`
	Salt4           uint64 `json:"salt4"`
	Salt5           uint64 `json:"salt5"`
	IsDisplayActive bool   `json:"isDisplayActive"`
	PopupDocFor     string `json:"popupDocFor"`
}
