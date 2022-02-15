package answer

import (
	"errors"
	"net/http"

	"github.com/ksaucedo002/answer/errores"
	"github.com/ksaucedo002/ctxman"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

const (
	response_data    = "data"
	response_message = "info"
	response_error   = "error"
)

type ResponseDetails struct {
	NumItems int    `json:"num_items"`
	OffSet   int    `json:"offset"`
	Limit    int    `json:"limit,omitempty"`
	Omits    string `json:"omits,omitempty"`
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
func OKDetails(c echo.Context, ctx ctxman.Ctxx, payload interface{}, num_elements int) error {
	params := ctx.GetParams()
	return c.JSON(http.StatusOK, &Response{
		Type: response_data,
		Data: payload,
		Details: &ResponseDetails{
			OffSet:   params.OffSet(),
			Limit:    params.Limit(),
			Omits:    params.Omitfiels(),
			NumItems: num_elements,
		},
	})
}
func Message(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, &Response{
		Type:    response_message,
		Message: message,
	})
}

func ErrorResponse(c echo.Context, err error) error {
	var errc *errores.CustomError
	code := 400
	message := "algo paso, hubo un error no esperado"
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
	return c.JSON(code, &Response{Type: response_error, Message: message})
}
func JSONErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequestf(nil, errores.ErrInvalidJSON))
}
func QueryErrorResponse(c echo.Context) error {
	return ErrorResponse(c, errores.NewBadRequestf(nil, errores.ErrInvalidQueryParam))
}
