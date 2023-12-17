package helper

import (
	"regexp"
	"strings"

	"github.com/dhimweray222/test-BE-uninet/exception"
)

func IsEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(e)
}

func ValidateStruct(err error) error {

	if strings.Contains(err.Error(), "Name") && strings.Contains(err.Error(), "required") {
		return exception.ErrorBadRequest("Name required.")
	}
	if strings.Contains(err.Error(), "Email") && strings.Contains(err.Error(), "required") {
		return exception.ErrorBadRequest("Email required.")
	}
	if strings.Contains(err.Error(), "Password") && strings.Contains(err.Error(), "required") {
		return exception.ErrorBadRequest("Password required.")
	}
	if strings.Contains(err.Error(), "Phone") && strings.Contains(err.Error(), "required") {
		return exception.ErrorBadRequest("Phone required.")
	}
	return exception.ErrorBadRequest(err.Error())
}
