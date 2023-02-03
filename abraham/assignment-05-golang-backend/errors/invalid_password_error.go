package errors

type InvalidPasswordError struct {
	// empty
}

func (e InvalidPasswordError) Error() string {
	return "invalid password"
}
