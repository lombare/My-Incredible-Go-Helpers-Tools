//
// filename:  errors.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irs

import (
	"fmt"
)

type HttpError struct {
	HttpCode int    `json:"httpCode"`
	ErrCode  Code   `json:"errCode"`
	Message  string `json:"message"`
}

func (n HttpError) Error() string {
	return fmt.Sprint("httpCode:", n.HttpCode, ". message:", n.Message)
}

func MakeHttpError(httpCode int, message ...interface{}) error {
	return HttpError{
		HttpCode: httpCode,
		Message:  fmt.Sprint(message...),
	}
}

type ServerError struct {
	HttpError
	Err error `json:"err"`
}

func (n ServerError) Error() string {
	if n.Err != nil {
		return fmt.Sprint("httpCode:", n.HttpCode, ". message:", n.Message, ". error:", n.Err.Error())
	} else {
		return fmt.Sprint("httpCode:", n.HttpCode, ". message:", n.Message, ". error: unknown")
	}
}

func MakeServerError(httpCode int, err error, message ...interface{}) error {
	if len(message) == 0 {
		message = []interface{}{
			Code(httpCode).Message(),
		}
	}
	return ServerError{
		HttpError: HttpError{
			HttpCode: httpCode,
			Message:  fmt.Sprint(message...),
		},
		Err: err,
	}
}

func MakeCodeError(code Code, err ...error) error {
	status, ok := ResponseStatuses[code]
	if !ok {
		return fmt.Errorf("unknown status code '%v'", code)
	}

	if err != nil && len(err) >= 1 {
		return MakeServerError(status.HttpCode, err[0], status.Message)
	}
	return MakeHttpError(status.HttpCode, status.Message)
}
