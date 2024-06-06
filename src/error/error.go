package custom_error

type Error interface {
	error
	GetStatusCode() int
}

type CustomError struct {
	statusCode int
	errMessage string
}
