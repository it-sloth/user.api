package error

type InternalError interface {
	// Redable message
	Error() string
	// Random string like "F82AP0L9"
	Code() string
	// HTTP status like 200 (OK)
	Status() int
}

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

// Make creates a new InternalError with the specified message, code, and HTTP status.
func Make(message string, code string, status int) InternalError {
	return &HttpError{
		message: message,
		code:    code,
		status:  status,
	}
}