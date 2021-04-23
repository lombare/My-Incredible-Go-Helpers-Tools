//
// filename:  error.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irsa

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	irs "github.com/lombare/might/sender"
)

type errorResponseTemplate struct {
	ErrCode    irs.Code `json:"errCode"`
	Message    string   `json:"message"`
}

func SendError(c echo.Context, err error) error {
	id := uuid.NewString()

	e := irs.HttpError{
		HttpCode: http.StatusInternalServerError,
		ErrCode:  0,
		Message:  http.StatusText(http.StatusInternalServerError),
	}

	if he, ok := err.(*echo.HTTPError); ok {
		code := irs.Code(he.Code)
		e.HttpCode = code.HttpCode()
		e.ErrCode = code
		e.Message = code.Message()
	} else if se, ok := err.(*irs.ServerError); ok {
		c.Logger().Error("id="+id, fmt.Sprint(se.Err))
		c.Response().Header().Set("X-Error-Id", id)
		e = se.HttpError
	} else if se, ok := err.(irs.ServerError); ok {
		c.Logger().Error("id="+id, fmt.Sprint(se.Err))
		c.Response().Header().Set("X-Error-Id", id)
		e = se.HttpError
	} else if he, ok := err.(*irs.HttpError); ok {
		e = *he
	} else if he, ok := err.(irs.HttpError); ok {
		e = he
	} else {
		c.Logger().Error("id="+id, fmt.Sprint(e))
		c.Response().Header().Set("X-Error-Id", id)
	}

	return RawSend(c, e.HttpCode, errorResponseTemplate{
		ErrCode:    e.ErrCode,
		Message:    e.Message,
	})
}

func ErrorHandler(err error, ctx echo.Context) {
	e := SendError(ctx, err)
	if e != nil {
		ctx.Logger().Fatal(err)
	}
}
