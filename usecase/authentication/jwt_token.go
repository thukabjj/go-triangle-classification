package authentication

import "github.com/thukabjj/go-triangle-classification/domain/authentication"

type JwtToken interface {
	ValidateCredentials(username string, password string) bool
	GenerateToken(username string, password string) (*authentication.Authentication, error)
}
