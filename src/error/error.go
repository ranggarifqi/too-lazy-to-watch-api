package custom_error

type Error interface {
	error
	GetStatusCode() int
}
