package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thukabjj/go-triangle-classification/domain"
	"golang.org/x/crypto/bcrypt"
)

type JwtTokenImpl struct{}

var USERNAME = "triangle"
var PASSWORD = generatedHashPassword("classification")
var EXPIRATION_TIME = time.Now().Add(time.Minute * 5).Unix()

func (j *JwtTokenImpl) ValidateCredentials(username string, password string) bool {
	return username == USERNAME && compareHashAndPassword(PASSWORD, password)
}

func (j *JwtTokenImpl) GenerateToken(username string, password string) (*domain.Authentication, error) {
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

	authenticatedToken := &domain.Authentication{
		Username:       USERNAME,
		Token:          tokenString,
		Type:           domain.AuthenticationTypeBearer,
		ExpirationTime: EXPIRATION_TIME,
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
