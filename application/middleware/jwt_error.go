package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thukabjj/go-triangle-classification/usecase/authentication/entity"
)

func JwtValidationsError(c *gin.Context) {

	c.Next()
	if len(c.Errors) > 0 {
		var response Error

		for _, err := range c.Errors {
			var na entity.NotAuthorizedError

			if errors.Is(err.Err, &na) {
				response.Code = http.StatusUnauthorized
				response.Status = http.StatusText(http.StatusUnauthorized)
				response.Details = map[string]string{"error": err.Err.Error()}
			}
			c.JSON(response.Code, response)
		}
	}
}
