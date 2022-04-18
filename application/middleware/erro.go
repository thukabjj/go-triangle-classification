package middleware

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Code    int               `json:"code"`
	Status  string            `json:"status"`
	Datails map[string]string `json:"datails"`
}

type Errors struct {
	Errors []Error `json:"errors"`
}

func ErrorHandler(c *gin.Context) {
	c.Next()

	var response Error

	for _, err := range c.Errors {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {

			response.Code = http.StatusUnprocessableEntity
			response.Status = http.StatusText(http.StatusBadRequest)
			me := make(map[string]string)
			for _, fe := range ve {
				me[fe.Field()] = getErrorMsg(fe)
			}
			response.Datails = me
			c.JSON(response.Code, response)

		}

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
