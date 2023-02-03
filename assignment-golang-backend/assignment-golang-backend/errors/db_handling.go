package errors

import "gorm.io/gorm"

func GetErrorByDB(err error, msg string) error {
	switch err {
	case gorm.ErrRecordNotFound:
		return ErrRecordNotFound(msg)
	default:
		return ErrInternalServer(err.Error())
	}

}
