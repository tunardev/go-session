package utils

import "strings"

type JError struct {
	Message string `json:"message"`
	Success bool `json:"success"`
}

func NewError(err error, success bool) JError {
	jerr := JError{
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