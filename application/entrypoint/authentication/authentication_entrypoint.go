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

// PostAuthentication             godoc
// @Summary      Make the authentication of an username
// @Description  Takes the username and the password from the Header and valid this information. Return JSON with the JWT information.
// @Tags         authentication
// @Produce      json
// @Param        username  header      string  true  "username"
// @Param        password  header      string  true  "password"
// @Success      201   {object}  entity.AuthenticationEntrypointResponse "JWT informations"
// @Failure      401   {object}  middleware.Error "User not authorized!"
// @Router       /auth/login [post]
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
