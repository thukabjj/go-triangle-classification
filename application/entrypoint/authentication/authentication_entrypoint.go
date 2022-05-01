package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/application/entrypoint/authentication/entity"
	"github.com/thukabjj/go-triangle-classification/usecase/authentication"
	AuthenticationUseCaseEntity "github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
)

type AuthenticationEntrypoint interface {
	Login(ctx *gin.Context)
}

type AuthenticationEntrypointImpl struct {
	AuthenticationUseCase authentication.AuthenticationUseCase
}

func (e *AuthenticationEntrypointImpl) Login(ctx *gin.Context) {

	username := ctx.Request.Header.Get("username")
	password := ctx.Request.Header.Get("password")

	if username == "" || password == "" {
		ctx.Error(&AuthenticationUseCaseEntity.NotAuthorizedError{})
		return
	}

	authenticatedToken, err := e.AuthenticationUseCase.Authenticate(username, password)

	if err != nil {
		ctx.Error(&AuthenticationUseCaseEntity.NotAuthorizedError{})
		return
	}

	ctx.JSON(201, &entity.AuthenticationEntrypointResponse{
		Username:       authenticatedToken.Username,
		Token:          authenticatedToken.Token,
		Type:           string(authenticatedToken.Type),
		ExpirationTime: authenticatedToken.ExpirationTime,
	})

}
