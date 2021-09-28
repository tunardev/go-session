package utils

import (
	"strings"

	"github.com/tunardev/go-session/models"
)

func NewError(err error, success bool) models.Error {
	newError := models.Error{
		Message: "generic error",
		Success: false,
	}
	if err != nil {
		newError.Message = err.Error()
		newError.Success = success 
	}
	return newError
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}