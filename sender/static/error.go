//
// filename:  error.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irss

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	irs "github.com/lombare/might/sender"
)

func SendError(c echo.Context, err error) error {
	id := uuid.NewString()

	if e, ok := err.(*irs.ServerError); ok {
		c.Logger().Error("id="+id, fmt.Sprint(e.Err))
		c.Response().Header().Set("X-Error-Id", id)
		return Send(c, e.HttpCode, nil, e.Message)
	} else if e, ok := err.(*irs.HttpError); ok {
		return Send(c, e.HttpCode, nil, e.Message)
	} else {
		c.Logger().Error("id="+id, fmt.Sprint(e))
		c.Response().Header().Set("X-Error-Id", id)
		return Send(c, http.StatusInternalServerError, nil, http.StatusText(http.StatusInternalServerError))
	}
}
