package response

import (
	"net/http"
)

var messageToHTTPCode = map[error]int{
	ErrBadRequest:      http.StatusBadRequest,
	ErrUnauthenticated: http.StatusForbidden,
	ErrNotFound:        http.StatusNotFound,
	ErrInternalServer:  http.StatusInternalServerError,
	ErrTimeoutError:    http.StatusGatewayTimeout,
}

//GetHTTPCode will get http code based on error message
func GetHTTPCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	httpCode, ok := messageToHTTPCode[err]
	if !ok {
		httpCode = http.StatusInternalServerError
	}
	return httpCode
}
