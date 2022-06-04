package authentication

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/thukabjj/go-triangle-classification/domain"
	"github.com/thukabjj/go-triangle-classification/mocks"
	AuthenticationUseCaseEntity "github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
)

func Test_AuthenticationEntrypointImpl(t *testing.T) {
	//Given
	mockerdAuthenticationUseCase := mocks.NewMockAuthenticationUseCase(gomock.NewController(t))

	expectedResponse := &domain.Authentication{
		Username:       "xpto",
		Token:          "dasas-asdasd-asdsa-saasd",
		Type:           domain.AuthenticationTypeBearer,
		ExpirationTime: 3600,
	}
	authenticationEntrypoint := &AuthenticationEntrypointImpl{
		AuthenticationUseCase: mockerdAuthenticationUseCase,
	}

	//When
	mockerdAuthenticationUseCase.EXPECT().Authenticate(gomock.Eq("username"), gomock.Eq("password")).Return(expectedResponse, nil)

	//Then
	r := gin.Default()
	req, _ := http.NewRequest(http.MethodPost, "/auth/login", nil)
	req.Header.Set("username", "username")
	req.Header.Set("password", "password")
	rr := httptest.NewRecorder()
	r.POST("/auth/login", authenticationEntrypoint.Login)
	r.ServeHTTP(rr, req)

	var response domain.Authentication
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.NotNil(t, response)
	assert.EqualValues(t, http.StatusCreated, rr.Code)
	assert.EqualValues(t, "xpto", response.Username)
	assert.EqualValues(t, "dasas-asdasd-asdsa-saasd", response.Token)
	assert.EqualValues(t, domain.AuthenticationTypeBearer, response.Type)
	assert.EqualValues(t, 3600, response.ExpirationTime)
}

func TestAuthenticationEntrypointImpl_Login(t *testing.T) {

	//Given
	mockerdAuthenticationUseCase := mocks.NewMockAuthenticationUseCase(gomock.NewController(t))

	expectedResponse := &domain.Authentication{
		Username:       "xpto",
		Token:          "dasas-asdasd-asdsa-saasd",
		Type:           domain.AuthenticationTypeBearer,
		ExpirationTime: 3600,
	}

	expectedError := &AuthenticationUseCaseEntity.NotAuthorizedError{}

	authenticationEntrypoint := &AuthenticationEntrypointImpl{
		AuthenticationUseCase: mockerdAuthenticationUseCase,
	}

	//When
	mockerdAuthenticationUseCase.EXPECT().Authenticate(gomock.Eq("username"), gomock.Eq("password")).Return(expectedResponse, nil)
	mockerdAuthenticationUseCase.EXPECT().Authenticate(gomock.Eq("username1"), gomock.Eq("password1")).Return(nil, expectedError)

	//Then
	type args struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		e       *AuthenticationEntrypointImpl
		args    args
		want    *domain.Authentication
		wantErr bool
	}{
		{
			name: "Should return status code 201 with authentication token",
			e:    authenticationEntrypoint,
			args: args{
				username: "username",
				password: "password",
			},
			want:    expectedResponse,
			wantErr: false,
		},
		{
			name: "Should return status code 401 with error",
			e:    authenticationEntrypoint,
			args: args{
				username: "username1",
				password: "password1",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := gin.Default()
			req, _ := http.NewRequest(http.MethodPost, "/auth/login", nil)
			req.Header.Set("username", tt.args.username)
			req.Header.Set("password", tt.args.password)
			rr := httptest.NewRecorder()
			r.POST("/auth/login", authenticationEntrypoint.Login)
			r.ServeHTTP(rr, req)

			var response *domain.Authentication
			err := json.Unmarshal(rr.Body.Bytes(), &response)

			if (err != nil) != tt.wantErr {
				t.Errorf("AuthenticationEntrypoint.Login error = %v, wantErr %v", err, tt.wantErr)
				return

			}

			if !reflect.DeepEqual(response, tt.want) {
				t.Errorf("AuthenticationUseCaseImpl.Authenticate() = %v, want %v", &response, tt.want)
			}

		})
	}
}
