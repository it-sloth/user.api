package error

import "net/http"

type Internal struct {
	message string
	code    string
	status  int
}

func (e *Internal) Error() string {
	return e.message
}

func (e *Internal) Code() string {
	return e.code
}

func (e *Internal) Status() int {
	return e.status
}

func NewInternal(message string, code string) InternalError {
	return &Internal{
		message: message,
		code:    code,
		status:  http.StatusInternalServerError,
	}
}
