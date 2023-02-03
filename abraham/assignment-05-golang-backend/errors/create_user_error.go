package errors

type CreateUserError struct {

}

func (e CreateUserError) Error() string {
	return "create user error"
}