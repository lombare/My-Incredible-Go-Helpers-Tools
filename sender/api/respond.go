//
// filename:  respond.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irsa

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type responseTemplate struct {
	Content    interface{} `json:"content"`
}

func Send(c echo.Context, httpCode int, payload interface{}) error {
	if id, ok := c.Get("request.id").(uuid.UUID); ok {
		c.Response().Header().Set("X-Request-ID", id.String())
	}

	return RawSend(c, httpCode, responseTemplate{
		Content:    payload,
	})
}

func RawSend(c echo.Context, httpCode int, data interface{}) error {
	switch c.Request().Header.Get("accept") {
	case "application/xml":
		fallthrough
	case "text/xml":
		return c.XML(httpCode, data)
	case "application/json":
		fallthrough
	default:
		return c.JSON(httpCode, data)
	}
}
