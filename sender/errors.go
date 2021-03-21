package sender

import (
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

func AddStatus(httpCode, code int, message string) {
	ResponseStatuses[httpCode] = Status{
		code,
		message,
	}
}
