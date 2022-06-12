package authorization

import "github.com/thukabjj/go-triangle-classification/usecase/gateway/jwt"

type AuthorizationUseCase interface {
	IsAuthorized(token string) bool
}

type AuthorizationUseCaseImpl struct {
	JwtToken jwt.JwtToken
}

func (c *AuthorizationUseCaseImpl) IsAuthorized(token string) bool {
	return c.JwtToken.ValidateToken(token)
}
