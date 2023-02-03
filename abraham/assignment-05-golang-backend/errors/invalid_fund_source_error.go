package errors 

type InvalidFundSourceError struct {
	// empty
}

func (e InvalidFundSourceError) Error() string {
	return "invalid fund source"
}