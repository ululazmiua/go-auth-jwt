package exception

type CustomBadRequestError struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e CustomBadRequestError) Error() string {
	return e.error
}

func NewCustomBadRequestError(error string) CustomBadRequestError {
	return CustomBadRequestError{error: error}
}
