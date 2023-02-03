package helpers

type JsonResponse struct {
	Data         interface{}
	StatusCode   int
	Message      string
	Error        bool
	ErrorMessage string
}
