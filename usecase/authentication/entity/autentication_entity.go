package entity

type AuthenticationType string

const (
	AuthenticationTypeBearer AuthenticationType = "Bearer"
)

type AuthenticationEntity struct {
	Username       string
	Token          string
	Type           AuthenticationType
	ExpirationTime int64
}
