//
// filename:  with_stats.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package middlewares

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func WithStats() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("request.id", uuid.New())
			return next(c)
		}
	}
}
