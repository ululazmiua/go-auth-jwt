package exception

type CustomInternalServerError struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e CustomInternalServerError) Error() string {
	return e.error
}

func NewCustomInternalServerError(error string) CustomInternalServerError {
	return CustomInternalServerError{error: error}
}
