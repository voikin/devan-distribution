package errs

type ErrorIncorrectPassword struct {
	Msg string
}

func (e *ErrorIncorrectPassword) Error() string {
	return e.Msg
}

func NewErrorIncorrectPassword() error {
	return &ErrorIncorrectPassword{"incorrect password"}
}
