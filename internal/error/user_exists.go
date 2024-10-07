package error

import "net/http"

type UserExists struct {
	message string
	code    string
	status  int
}

func (e *UserExists) Error() string {
	return e.message
}

func (e *UserExists) Code() string {
	return e.code
}

func (e *UserExists) Status() int {
	return e.status
}

func NewUserExists(message string, code string) InternalError {
	return &UserExists{
		message: message,
		code:    code,
		status:  http.StatusConflict,
	}
}
