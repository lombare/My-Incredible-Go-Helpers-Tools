package irs

import (
	"fmt"
	"net/http"
)

const (
	Ok      = 0
	Created = 1
	Updated = 2
	Deleted = 3

	BadRequest              = 100
	NotSupportedAccept      = 101
	NotSupportedContentType = 102

	StatusPadding = 1000
)

type Status struct {
	HttpCode int
	Message  string
}

var ResponseStatuses = map[int]Status{
	Ok:      {http.StatusOK, "SUCCESS/OK"},
	Created: {http.StatusCreated, "SUCCESS/CREATED"},
	Updated: {http.StatusOK, "SUCCESS/UPDATED"},
	Deleted: {http.StatusNoContent, "SUCCESS/DELETED"},

	BadRequest:              {http.StatusBadRequest, "ERRORS/REQUEST/BAD_REQUEST"},
	NotSupportedAccept:      {http.StatusNotAcceptable, "ERRORS/REQUEST/NOT_SUPPORTED_ACCEPT"},
	NotSupportedContentType: {http.StatusNotAcceptable, "ERRORS/REQUEST/NOT_SUPPORTED_CONTENT_TYPE"},
}

func AddStatus(code, httpCode int, message string) {
	ResponseStatuses[code] = Status{
		httpCode,
		message,
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
