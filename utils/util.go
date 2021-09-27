package utils

import "strings"

type JError struct {
	Error string `json:"error"`
	Success bool `json:"success"`
}

func NewError(err error) JError {
	jerr := JError{
		Error: "generic error",
		Success: false,
	}
	if err != nil {
		jerr.Error = err.Error()
	}
	return jerr
}

func NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}