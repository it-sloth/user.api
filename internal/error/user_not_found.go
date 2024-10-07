package error

import "net/http"

type UserNotFoundError struct {
	message string
	code    string
	status  int
}

func (e *UserNotFoundError) Error() string {
	return e.message
}

func (e *UserNotFoundError) Code() string {
	return e.code
}

func (e *UserNotFoundError) Status() int {
	return e.status
}

func NewUserNotFoundError(message string, code string) InternalError {
	return &UserNotFoundError{
		message: message,
		code:    code,
		status:  http.StatusNotFound,
	}
}
