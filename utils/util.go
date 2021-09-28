package utils

import "strings"

type JError struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}

func NewError(err error) JError {
	jerr := JError{
		Message: "generic error",
		Success: false,
	}
	if err != nil {
		jerr.Message = err.Error()
	}
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}