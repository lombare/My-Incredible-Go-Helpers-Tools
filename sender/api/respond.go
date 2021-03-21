package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"

	"github.com/lombare/might/sender"
)

type responseTemplate struct {
	HttpCode   int           `json:"httpCode"`
	HttpStatus string        `json:"httpStatus"`
	Message    string        `json:"message"`
	Content    interface{}   `json:"content"`
	Duration   time.Duration `json:"delay,omitempty"`
}

func SendJSON(c echo.Context, status int, payload interface{}, message ...interface{}) error {
	var d time.Duration
	if t, ok := c.Get("request.time").(time.Time); ok {
		d = time.Since(t)
	}

	return c.JSON(status, responseTemplate{
		HttpCode:   status,
		HttpStatus: http.StatusText(status),
		Message:    fmt.Sprint(message...),
		Content:    payload,
		Duration:   d,
	})
}

func SendXML(c echo.Context, status int, payload interface{}, message ...interface{}) error {
	var d time.Duration
	if t, ok := c.Get("request.time").(time.Time); ok {
		d = time.Since(t)
	}

	return c.XML(status, responseTemplate{
		HttpCode:   status,
		HttpStatus: http.StatusText(status),
		Message:    fmt.Sprint(message...),
		Content:    payload,
		Duration:   d,
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
		return SendXML(c, status, payload, message)
	case "application/json":
		fallthrough
	default:
		return SendJSON(c, status, payload, message)
	}
}

func SendCode(c echo.Context, code int, payload interface{}) error {
	status, ok := sender.ResponseStatuses[code]
	if !ok {
		return fmt.Errorf("unknown status code '%v'", code)
	}

	switch c.Request().Header.Get("accept") {
	case "application/xml":
		fallthrough
	case "text/xml":
		return SendXML(c, status.HttpCode, payload, status.Message)
	case "application/json":
		fallthrough
	default:
		return SendJSON(c, status.HttpCode, payload, status.Message)
	}
}
