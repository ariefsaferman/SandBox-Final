package errors

import (
	"errors"
	"fmt"
)

var ErrRecordNotFound = func(msg string) error {
	return errors.New(fmt.Sprintf("there is no record %s", msg))
}

var ErrInternalServer = func(msg string) error {
	return errors.New(fmt.Sprintf("internal server error :%s", msg))
}

var ErrInvalidPassword = errors.New("invalid password")
var ErrUserAlreadyRegister = errors.New("email is already used")
var ErrBadRequest = errors.New("bad request")
var ErrInvalidUser = errors.New("there is no user with that email")
var ErrTopUp = errors.New("error top up")
var ErrTransfer = errors.New("error transfer")
var ErrTargetWalletNotFound = errors.New("target wallet not found")
var ErrInsufficientBalance = errors.New("insufficient balance")
var ErrServerError = errors.New("internal server error")
var ErrListTransaction = errors.New("there is no transaction with that id")
var ErrUserNotFound = errors.New("there is no user with that id")
