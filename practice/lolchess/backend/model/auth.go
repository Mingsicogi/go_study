package model

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Realm       string `json:"realm"`
	TokenType   string `json:"token_type"`
}
