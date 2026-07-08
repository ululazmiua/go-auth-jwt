package exception

type NotFoundError struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e NotFoundError) Error() string {
	return e.error
}
func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{error: error}
}
