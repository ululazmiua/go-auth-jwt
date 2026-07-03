package exception

type EmailAlreadyExistsError struct {
	error string
}

func NewEmailAlreadyExistsError(error string) EmailAlreadyExistsError {
	return EmailAlreadyExistsError{error: error}
}
