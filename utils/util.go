package utils

import (
	"strings"

	"github.com/tunardev/go-session/models"
)

func NewError(err error, success bool) models.Error {
	jerr := models.Error{
		Message: "generic error",
		Success: false,
	}
	if err != nil {
		jerr.Message = err.Error()
		jerr.Success = success 
	}
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}