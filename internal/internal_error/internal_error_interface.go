package internal_error

type InternalErrorInterface interface {
	Error() string
	Code() string
}
