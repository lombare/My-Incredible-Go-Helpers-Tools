package irsa

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/lombare/might/sender"
)

type responseTemplate struct {
	HttpCode   int         `json:"httpCode"`
	HttpStatus string      `json:"httpStatus"`
	Message    string      `json:"message"`
	Content    interface{} `json:"content"`
}

func SendJSON(c echo.Context, status int, payload interface{}, message ...interface{}) error {
	return c.JSON(status, responseTemplate{
		HttpCode:   status,
		HttpStatus: http.StatusText(status),
		Message:    fmt.Sprint(message...),
		Content:    payload,
	})
}

func SendXML(c echo.Context, status int, payload interface{}, message ...interface{}) error {
	return c.XML(status, responseTemplate{
		HttpCode:   status,
		HttpStatus: http.StatusText(status),
		Message:    fmt.Sprint(message...),
		Content:    payload,
	})
}

func Send(c echo.Context, status int, payload interface{}, message ...interface{}) error {
	if id, ok := c.Get("request.id").(uuid.UUID); ok {
		c.Response().Header().Set("X-Request-ID", id.String())
	}

	switch c.Request().Header.Get("accept") {
	case "application/xml":
		fallthrough
	case "text/xml":
		return SendXML(c, status, payload, message...)
	case "application/json":
		fallthrough
	default:
		return SendJSON(c, status, payload, message...)
	}
}

func SendCode(c echo.Context, code int, content ...interface{}) error {
	status, ok := irs.ResponseStatuses[code]
	if !ok {
		return fmt.Errorf("unknown status code '%v'", code)
	}

	var p interface{}
	if len(content) != 0 {
		p = content[0]
	}
	switch c.Request().Header.Get("accept") {
	case "application/xml":
		fallthrough
	case "text/xml":
		return SendXML(c, status.HttpCode, p, status.Message)
	case "application/json":
		fallthrough
	default:
		return SendJSON(c, status.HttpCode, p, status.Message)
	}
}
