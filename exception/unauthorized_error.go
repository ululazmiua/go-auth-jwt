package exception

type UnauthorizedError struct {
	error string
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{error: error}
}
