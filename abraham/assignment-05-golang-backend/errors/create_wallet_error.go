package errors

type CreateWalletError struct {

}

func (e CreateWalletError) Error() string {
	return "create wallet error"
}