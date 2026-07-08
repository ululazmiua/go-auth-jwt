package exception

type EmailAlreadyExistsError struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e EmailAlreadyExistsError) Error() string {
	return e.error
}

func NewEmailAlreadyExistsError(error string) EmailAlreadyExistsError {
	return EmailAlreadyExistsError{error: error}
}
