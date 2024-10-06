package internal_error

type UserNotFoundError struct {
	message string
	code    string
}

func (e *UserNotFoundError) Error() string {
	return e.message
}

func (e *UserNotFoundError) Code() string {
	return e.code
}

func NewUserNotFoundError(message string, code string) InternalErrorInterface {
	return &UserNotFoundError{
		message: message,
		code:    code,
	}
}
