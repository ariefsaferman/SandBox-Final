package errors

type AmountOutOfRangeError struct {

}

func (e AmountOutOfRangeError) Error() string {
	return "amount out of range"
}