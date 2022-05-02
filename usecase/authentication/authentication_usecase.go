package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
	"golang.org/x/crypto/bcrypt"
)

var USERNAME = "triangle"
var PASSWORD = GeneratehashPassword("classification")
var EXPIRATION_TIME = time.Now().Add(time.Minute * 5).Unix()
var users = map[string]string{USERNAME: PASSWORD}

type AuthenticationUseCase interface {
	Authenticate(username string, password string) (*entity.AuthenticationEntity, error)
}

type AuthenticationUseCaseImpl struct{}

func (c *AuthenticationUseCaseImpl) Authenticate(username string, password string) (*entity.AuthenticationEntity, error) {

	if users[username] == "" || !CheckPasswordHash(password, users[username]) {
		return nil, &entity.NotAuthorizedError{}
	}

	var mySigningKey = []byte("secretkey")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = username
	claims["role"] = "ADMIN"
	claims["exp"] = EXPIRATION_TIME

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return nil, err
	}

	authenticatedToken := &entity.AuthenticationEntity{
		Username:       USERNAME,
		Token:          tokenString,
		Type:           entity.AuthenticationTypeBearer,
		ExpirationTime: EXPIRATION_TIME,
	}
	return authenticatedToken, nil

}

func GeneratehashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
