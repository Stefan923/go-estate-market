package response

import (
	error2 "github.com/Stefan923/go-estate-market/error"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	// User
	error2.EmailAlreadyUsed:   409,
	error2.RecordNotFound:     404,
	error2.InvalidCredentials: 401,
	error2.PermissionDenied:   403,
}

func TranslateErrorToStatusCode(err error) int {
	statusCode, success := StatusCodeMapping[err.Error()]
	if !success {
		return http.StatusInternalServerError
	}
	return statusCode
}
