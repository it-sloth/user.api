package error

type HttpError struct {
	message string
	code    string
	status  int
}

func (e *HttpError) Error() string {
	return e.message
}

func (e *HttpError) Code() string {
	return e.code
}

func (e *HttpError) Status() int {
	return e.status
}

func Make(message string, code string, status int) InternalError {
	return &HttpError{
		message: message,
		code:    code,
		status:  status,
	}
}
