//
// filename:  respond.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irss

import (
	"bufio"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func Send(c echo.Context, httpCode int, reader io.Reader, message ...interface{}) error {
	if id, ok := c.Get("request.id").(uuid.UUID); ok {
		c.Response().Header().Set("X-Request-ID", id.String())
	}
	c.Response().Header().Set("X-Message", fmt.Sprint(message...))

	// If there is no content then send nothing
	if reader == nil {
		return c.NoContent(httpCode)
	}

	// Wrap the io.Reader into a bufio.Reader. It allow to "unread" bytes.
	b := bufio.NewReader(reader)
	buff := make([]byte, 512)

	// Get the first 512 bytes to scan the file mime type
	if _, err := b.Read(buff); err != nil {
		return SendError(c, err)
	}

	// Detect the mime type with the 512 first bytes
	contentType := http.DetectContentType(buff)
	if err := b.UnreadByte(); err != nil {
		return SendError(c, err)
	}

	return c.Stream(httpCode, contentType, b)
}
