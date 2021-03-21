package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SendError(c echo.Context, err error) error {
	id := uuid.NewString()

	if e, ok := err.(*CriticalError); ok {
		c.Logger().Error("id="+id, fmt.Sprint(e.Err))
		c.Response().Header().Set("X-Error-Id", id)
		return Send(c, e.Code, nil, e.Message)
	} else if e, ok := err.(*NormalError); ok {
		return Send(c, e.Code, nil, e.Message)
	} else {
		c.Logger().Error("id="+id, fmt.Sprint(e))
		c.Response().Header().Set("X-Error-Id", id)
		return Send(c, http.StatusInternalServerError, nil, http.StatusText(http.StatusInternalServerError))
	}
}

type NormalError struct {
	Code    int
	Message string
}

func (n NormalError) Error() string {
	return fmt.Sprint("code:", n.Code, ". message:", n.Message)
}

func MakeNormalError(code int, message ...interface{}) error {
	return NormalError{
		Code:    code,
		Message: fmt.Sprint(message...),
	}
}

type CriticalError struct {
	NormalError
	Err error
}

func (n CriticalError) Error() string {
	return fmt.Sprint("code:", n.Code, ". message:", n.Message, ". error:", n.Error())
}

func MakeCriticalError(code int, err error, message ...interface{}) error {
	return CriticalError{
		NormalError: NormalError{
			Code:    code,
			Message: fmt.Sprint(message...),
		},
		Err: err,
	}
}
