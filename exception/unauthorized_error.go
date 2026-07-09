package exception

type UnauthorizedError struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e UnauthorizedError) Error() string {
	return e.error
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{error: error}
}
