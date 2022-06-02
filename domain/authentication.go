package domain

type AuthenticationType string

const (
	AuthenticationTypeBearer AuthenticationType = "Bearer"
)

type Authentication struct {
	Username       string
	Token          string
	Type           AuthenticationType
	ExpirationTime int64
}
