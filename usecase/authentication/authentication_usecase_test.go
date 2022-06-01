package authentication

import (
	"github.com/golang/mock/gomock"
	"github.com/thukabjj/go-triangle-classification/domain/authentication"
	"github.com/thukabjj/go-triangle-classification/mocks"
	"testing"

	"github.com/go-playground/assert/v2"
)

func Test_Authenticate(t *testing.T) {

	//given
	jwtTokenMocked := mocks.NewMockJwtToken(gomock.NewController(t))

	username := "triangle"
	password := "classification"

	AuthenticationUseCase := &AuthenticationUseCaseImpl{
		JwtToken: jwtTokenMocked,
	}

	response := &authentication.Authentication{
		Username:       "triangle",
		Token:          "klsnlksnklsnalsllaL",
		Type:           authentication.AuthenticationTypeBearer,
		ExpirationTime: 3600,
	}

	//when
	jwtTokenMocked.EXPECT().ValidateCredentials(gomock.Eq(username), gomock.Eq(password)).Return(true)
	jwtTokenMocked.EXPECT().GenerateToken(gomock.Eq(username), gomock.Eq(password)).Return(response, nil)

	//then
	result, err := AuthenticationUseCase.Authenticate(username, password)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Username, username)

}
