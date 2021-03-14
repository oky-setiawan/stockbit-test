package response

import "errors"

//Error Variabel
var (
	ErrBadRequest      = errors.New("Bad Request")
	ErrUnauthenticated = errors.New("Unauthorized")
	ErrNotFound        = errors.New("Not Found")
	ErrInternalServer  = errors.New("Internal Server Error")
	ErrTimeoutError    = errors.New("Timeout error")
)

//Error message
const (
	MsgBadRequest      ErrorMessage = "invalid request, please check your request param"
	MsgUnauthenticated ErrorMessage = "you don't have permission to access"
	MsgNotFound        ErrorMessage = "your request not found, please check is your data valid"
	MsgInternalServer  ErrorMessage = "internal server error"
)

type ErrorMessage string

func (e ErrorMessage) String() string {
	return string(e)
}

var errToMessage = map[error]ErrorMessage{
	ErrBadRequest:      MsgBadRequest,
	ErrUnauthenticated: MsgUnauthenticated,
	ErrNotFound:        MsgNotFound,
	ErrInternalServer:  MsgInternalServer,
}

//GetErrorMessage will get error message
func GetErrorMessage(err error) (msg ErrorMessage) {
	if err == nil {
		return
	}

	var ok bool
	msg, ok = errToMessage[err]
	if !ok {
		msg = MsgInternalServer
	}
	return
}
