package dto

type TokenResponse struct {
	AccessToken   string `json:"access_token"`
	ResponseToken string `json:"response_token"`
}
