package errors

type UserNotFoundError struct {
	// empty
}

func (e UserNotFoundError) Error() string {
	return "user not found"
}
