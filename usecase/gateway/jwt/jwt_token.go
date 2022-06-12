package jwt

import "github.com/thukabjj/go-triangle-classification/domain"

type JwtToken interface {
	ValidateCredentials(username string, password string) bool
	GenerateToken(username string, password string) (*domain.Authentication, error)
	ValidateToken(token string) bool
}
