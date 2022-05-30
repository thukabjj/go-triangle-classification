package authentication

import (
	"time"

	"github.com/thukabjj/go-triangle-classification/domain/authentication"
	"github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
	"golang.org/x/crypto/bcrypt"
)

var USERNAME = "triangle"
var PASSWORD = generatedHashPassword("classification")
var EXPIRATION_TIME = time.Now().Add(time.Minute * 5).Unix()

type AuthenticationUseCase interface {
	Authenticate(username string, password string) (*authentication.Authentication, error)
}

type AuthenticationUseCaseImpl struct {
	JwtToken JwtToken
}

func (c *AuthenticationUseCaseImpl) Authenticate(username string, password string) (*authentication.Authentication, error) {

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

func generatedHashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func compareHashAndPassword(passwordEncrypted string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordEncrypted), []byte(password)) == nil
}
