package authentication

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Authenticate(t *testing.T) {
	AuthenticationUseCase := &AuthenticationUseCaseImpl{}

	username := "triangle"
	password := "classification"

	result, err := AuthenticationUseCase.Authenticate(username, password)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Username, username)

}
