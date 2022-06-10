package answer

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/ksaucedo002/answer/errores"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	response_data    = "data"
	response_message = "info"
	response_error   = "error"
)

type ResponseDetails struct {
	NumItems int      `json:"num_items"`
	OffSet   int      `json:"offset"`
	Limit    int      `json:"limit"`
	Preloads []string `json:"preload,omitempty"`
}
type Response struct {
	Type    string           `json:"type,omitempty"` //error, response
	Message string           `json:"message,omitempty"`
	Data    interface{}      `json:"data,omitempty"`
	Details *ResponseDetails `json:"details,omitempty"`
}

func OK(c echo.Context, payload interface{}) error {
	return c.JSON(http.StatusOK, &Response{
		Type: response_data,
		Data: payload,
	})
}
func ResOK(payload interface{}) Response {
	return Response{
		Type: response_data,
		Data: payload,
	}
}
func payloadLen(payload interface{}) int {
	var vlen = 0
	switch reflect.TypeOf(payload).Kind() {
	case reflect.Slice:
		vlen = reflect.ValueOf(payload).Len()
	}
	return vlen
}
func OKDetails(c echo.Context, payload interface{}, p Params) error {
	return c.JSON(http.StatusOK, &Response{
		Type: response_data,
		Data: payload,
		Details: &ResponseDetails{
			OffSet:   p.GetOffSet(),
			Limit:    p.GetLimit(),
			NumItems: payloadLen(payload),
			Preloads: p.GetPreloads(),
		},
	})
}

func Message(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, &Response{
		Type:    response_message,
		Message: message,
	})
}
func SMS(message string) Response {
	return Response{
		Type:    response_message,
		Message: message,
	}
}
func unwrap(err error) (code int, message string) {
	var errc *errores.CustomError
	code = 400
	message = "algo paso, hubo un error no esperado"
	if errors.As(err, &errc) {
		code = errc.HttpCode
		message = errc.ErrorMessage
	}
	go func(e error, ec *errores.CustomError) {
		if ec == nil {
			logrus.Error(e.Error())
			return
		}
		if ec.GetError() != nil {
			logrus.Error(e.Error())
			return
		}
	}(err, errc)
	return code, message
}
func ErrorResponse(c echo.Context, err error) error {
	code, message := unwrap(err)
	return c.JSON(code, &Response{Type: response_error, Message: message})
}
func Error(err error) Response {
	_, message := unwrap(err)
	return Response{Type: response_error, Message: message}
}
func JSONErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequestf(nil, errores.ErrInvalidJSON))
}
func QueryErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequestf(nil, errores.ErrInvalidQueryParam))
}
