package exception

type InvalidEmailPassword struct {
	error string
}

// mengimplementasi interface error bawan golang
func (e InvalidEmailPassword) Error() string {
	return e.error
}

func NewInvalidEmailPassword(error string) InvalidEmailPassword {
	return InvalidEmailPassword{error: error}
}
