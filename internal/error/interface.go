package error

type InternalError interface {
	Error() string
	Code() string
	Status() int
}
