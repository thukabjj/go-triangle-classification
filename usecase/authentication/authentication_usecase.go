package authentication

import (
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
	"github.com/thukabjj/go-triangle-classification/usecase/gateway/jwt"
)

type AuthenticationUseCase interface {
	Authenticate(username string, password string) (*domain.Authentication, error)
}

type AuthenticationUseCaseImpl struct {
	JwtToken jwt.JwtToken
}

func (c *AuthenticationUseCaseImpl) Authenticate(username string, password string) (*domain.Authentication, error) {

	credentialsIsValid := c.JwtToken.ValidateCredentials(username, password)

	if !credentialsIsValid {
		return nil, &entity.NotAuthorizedError{}
	}

	authenticatedToken, err := c.JwtToken.GenerateToken(username, password)

	if err != nil {
		return nil, err
	}

	return authenticatedToken, nil

}
