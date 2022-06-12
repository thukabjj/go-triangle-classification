package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infraJwt "github.com/thukabjj/go-triangle-classification/infrastructure/jwt"
	"github.com/thukabjj/go-triangle-classification/usecase/authorization"
)

func ValidateToken(c *gin.Context) {
	authorizationUseCaseImpl := &authorization.AuthorizationUseCaseImpl{JwtToken: &infraJwt.JwtTokenImpl{}}

	token := c.GetHeader("Authorization")
	isAuthorized := authorizationUseCaseImpl.IsAuthorized(token)
	if token == "" || !isAuthorized {
		response := &Errors{
			Errors: []Error{
				Error{
					Code:   http.StatusUnauthorized,
					Status: http.StatusText(http.StatusUnauthorized),
					Details: map[string]string{
						"error": "User not authorized!",
					},
				},
			},
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	} else {
		c.Next()
	}

}
