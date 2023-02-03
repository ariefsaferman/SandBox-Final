package errors

import "errors"

const (
	ErrCodeDuplicate           = "DUPLICATE_RECORD"
	ErrCodeInternalServerError = "INTERNAL_SERVER_ERROR"
	ErrCodeBadRequest          = "BAD_REQUEST"
	ErrCodeUnauthorized        = "UNAUTHORIZED"
	ErrCodeForbidden           = "FORBIDDEN_ACCESS"
	ErrCodeRouteNotFound       = "ROUTE_NOT_FOUND"
)

var (
	ErrInvalidBody             = errors.New("invalid body request")
	ErrUserNotFound            = errors.New("user not found")
	ErrWrongPassword           = errors.New("password mismatch")
	ErrSourceOfFundNotFound    = errors.New("source of fund not found")
	ErrWalletNotFound          = errors.New("wallet not found")
	ErrRecipientWalletNotFound = errors.New("recipient wallet not found")
	ErrSenderWalletNotFound    = errors.New("sender wallet not found")
	ErrInsufficientBalance     = errors.New("insufficient balance")
	ErrRouteNotFound           = errors.New("the requested route is not exist")
	ErrFailedToHash            = errors.New("failed to hash")
	ErrFailedToGenerateToken   = errors.New("failed to generate token")
	ErrInternalServerError     = errors.New("internal server error")
)
