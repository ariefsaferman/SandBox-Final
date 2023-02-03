package errors

type TransactionError struct {

}

func (e TransactionError) Error() string {
	return "transaction error"
}