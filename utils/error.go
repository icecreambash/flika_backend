package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorMessages struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ExceptionError(ctx *gin.Context, err error) {
	var validatorsError validator.ValidationErrors
	var validatorUnmarshal *json.UnmarshalTypeError
	if errors.As(err, &validatorUnmarshal) {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": []ErrorMessages{
				{
					Field:   validatorUnmarshal.Field,
					Message: validatorUnmarshal.Type.Name(),
				},
			},
		})
		return
	} else if errors.As(err, &validatorsError) {
		response := make([]ErrorMessages, len(validatorsError))

		for i, msg := range validatorsError {
			response[i] = ErrorMessages{
				Field:   msg.Field(),
				Message: GetErrorMessage(msg),
			}
		}

		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": response,
		})
		return
	}

	ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"errors": err,
	})
}

func GetErrorMessage(error validator.FieldError) string {
	switch error.Tag() {
	case "required":
		return fmt.Sprintf("Field %s is required", error.Param())

	case "email":
		return "Must been email format"

	case "min":
		return fmt.Sprintf("Min value %s", error.Param())
	}

	return "ERROR NOT FOUND"
}
