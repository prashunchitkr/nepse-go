package models

type Prove struct {
	ServerTime      uint64 `json:"serverTime"`
	Salt            string `json:"salt"`
	AccessToken     string `json:"accessToken"`
	TokenType       string `json:"tokenType"`
	RefreshToken    string `json:"refreshToken"`
	Salt1           uint32 `json:"salt1"`
	Salt2           uint32 `json:"salt2"`
	Salt3           uint32 `json:"salt3"`
	Salt4           uint32 `json:"salt4"`
	Salt5           uint32 `json:"salt5"`
	IsDisplayActive bool   `json:"isDisplayActive"`
	PopupDocFor     string `json:"popupDocFor"`
}
