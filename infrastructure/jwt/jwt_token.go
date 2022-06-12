package jwt

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/thukabjj/go-triangle-classification/domain"
	"golang.org/x/crypto/bcrypt"
)

type JwtTokenImpl struct{}

type UniqueClaims struct {
	jwt.StandardClaims
	TokenId string `json:"jti,omitempty"`
}

var USERNAME = "triangle"
var PASSWORD = generatedHashPassword("classification")
var EXPIRATION_TIME = time.Now().Add(time.Minute * 5).Unix()
var MY_SECRET = []byte("secretkey")

func (j *JwtTokenImpl) ValidateCredentials(username string, password string) bool {
	return username == USERNAME && compareHashAndPassword(PASSWORD, password)
}

func (j *JwtTokenImpl) GenerateToken(username string, password string) (*domain.Authentication, error) {
	now := time.Now()
	bits := make([]byte, 12)
	_, err := rand.Read(bits)
	if err != nil {
		panic(err)
	}
	claims := UniqueClaims{
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: EXPIRATION_TIME,
			Issuer:    "telescope",
		},
		TokenId: base64.StdEncoding.EncodeToString(bits),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(MY_SECRET)

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

func (j *JwtTokenImpl) ValidateToken(token string) bool {

	tokenWithoutBearer := token[7:]

	tokenClaims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenWithoutBearer, tokenClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("There was an error in parsing")
		}
		return MY_SECRET, nil
	})

	if err != nil {
		return false
	}
	return true
}

func generatedHashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func compareHashAndPassword(passwordEncrypted string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(passwordEncrypted), []byte(password)) == nil
}
