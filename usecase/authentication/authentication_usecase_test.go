package authentication

import (
	"errors"
	"reflect"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/mocks"
)

func Test_Authenticate(t *testing.T) {

	//given
	jwtTokenMocked := mocks.NewMockJwtToken(gomock.NewController(t))

	username := "triangle"
	password := "classification"

	AuthenticationUseCase := &AuthenticationUseCaseImpl{
		JwtToken: jwtTokenMocked,
	}

	expected := &domain.Authentication{
		Username:       "triangle",
		Token:          "klsnlksnklsnalsllaL",
		Type:           domain.AuthenticationTypeBearer,
		ExpirationTime: 3600,
	}

	//when
	jwtTokenMocked.EXPECT().ValidateCredentials(gomock.Eq(username), gomock.Eq(password)).Return(true)
	jwtTokenMocked.EXPECT().GenerateToken(gomock.Eq(username), gomock.Eq(password)).Return(expected, nil)

	//then
	result, err := AuthenticationUseCase.Authenticate(username, password)
	assert.Equal(t, err, nil)
	assert.Equal(t, result.Username, username)
	assert.Equal(t, result.Token, expected.Token)
	assert.Equal(t, result.Type, expected.Type)
	assert.Equal(t, result.ExpirationTime, expected.ExpirationTime)

}

func TestAuthenticationUseCaseImpl_Authenticate(t *testing.T) {

	//given
	jwtTokenMocked := mocks.NewMockJwtToken(gomock.NewController(t))

	username := "triangle"
	password := "classification"

	AuthenticationUseCase := &AuthenticationUseCaseImpl{
		JwtToken: jwtTokenMocked,
	}

	expected := &domain.Authentication{
		Username:       "triangle",
		Token:          "klsnlksnklsnalsllaL",
		Type:           domain.AuthenticationTypeBearer,
		ExpirationTime: 3600,
	}

	//when
	jwtTokenMocked.EXPECT().ValidateCredentials(gomock.Eq("triangle"), gomock.Eq("classification")).Return(true)
	jwtTokenMocked.EXPECT().ValidateCredentials(gomock.Eq("other2"), gomock.Eq("other2")).Return(true)

	jwtTokenMocked.EXPECT().ValidateCredentials(gomock.Eq("other"), gomock.Eq("other")).Return(false)

	jwtTokenMocked.EXPECT().GenerateToken(gomock.Eq("triangle"), gomock.Eq("classification")).Return(expected, nil)

	jwtTokenMocked.EXPECT().GenerateToken(gomock.Eq("other2"), gomock.Eq("other2")).Return(nil, errors.New("Something went wrong"))

	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		c       *AuthenticationUseCaseImpl
		args    args
		want    *domain.Authentication
		wantErr bool
	}{
		{"Should return JWT Token with success!", AuthenticationUseCase, args{username, password}, expected, false},
		{"Should return	error when credentials are invalid!", AuthenticationUseCase, args{"other", "other"}, nil, true},
		{"Should return error whe JWT Token generation fails!", AuthenticationUseCase, args{"other2", "other2"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Authenticate(tt.args.username, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticationUseCaseImpl.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthenticationUseCaseImpl.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
