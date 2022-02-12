package errores

import (
	"fmt"
	"net/http"
)

type CustomError struct {
	HttpCode     int
	Err          error
	ErrorMessage string
}

func newErrf(err error, message string, httpcode int) error {
	return &CustomError{
		Err:          err,
		ErrorMessage: message,
		HttpCode:     httpcode,
	}
}

func NewBadRequestf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusBadRequest)
}
func NewInternalf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusInternalServerError)
}
func NewUnsupported(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusUnsupportedMediaType)
}
func NewUnauthorizedf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusUnauthorized)
}
func NewForbiddenf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusForbidden)
}
func NewNotFoundf(err error, format string, a ...interface{}) error {
	return newErrf(err, fmt.Sprintf(format, a...), http.StatusNotFound)
}
func (e *CustomError) Error() string {
	if e.Err == nil {
		return "error trivial "
	}
	return e.Err.Error()
}
func (e *CustomError) GetError() error {
	return e.Err
}
func (e *CustomError) Message() string {
	return e.ErrorMessage
}
