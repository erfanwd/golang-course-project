package helpers

import (
	"net/http"

	"github.com/erfanwd/golang-course-project/pkg/service_errors"
)

var StatusCodeMapping = map[string]int{
	// OPT
	service_errors.OtpExists:     409,
	service_errors.OtpIsNotValid: 400,
	service_errors.OtpUsed:       409,

}

func TranslateErrorToStatusCode(err error) int {
	value, ok := StatusCodeMapping[err.Error()]
	if !ok {
		return http.StatusInternalServerError
	}
	return value
}
