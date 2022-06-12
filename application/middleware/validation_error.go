package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func ValidationErrorsMiddleware(c *gin.Context) {
	c.Next()

	if len(c.Errors) > 0 {
		var response Error
		for _, err := range c.Errors {
			var ve validator.ValidationErrors

			if errors.As(err, &ve) {

				response.Code = http.StatusUnprocessableEntity
				response.Status = http.StatusText(http.StatusUnprocessableEntity)
				me := make(map[string]string)
				for _, fe := range ve {
					me[fe.Field()] = getErrorMsg(fe)
				}
				response.Details = me

			}
		}
		c.JSON(response.Code, response)
	}
}

func getErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "gt":
		return "Should be greater than " + fe.Param()
	}
	return "Unknown error"
}
