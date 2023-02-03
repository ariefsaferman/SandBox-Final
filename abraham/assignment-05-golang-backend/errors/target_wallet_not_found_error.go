package errors

type TargetWalletNotFoundError struct {

}

func (e TargetWalletNotFoundError) Error() string {
	return "target wallet not found"
}