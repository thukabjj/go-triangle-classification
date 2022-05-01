package entity

type AuthenticationEntrypointResponse struct {
	Username       string `json:"username"`
	Token          string `json:"token"`
	Type           string `json:"type"`
	ExpirationTime int64  `json:"expirationTime"`
}
