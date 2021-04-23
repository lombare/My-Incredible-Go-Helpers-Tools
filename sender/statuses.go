//
// filename:  statuses.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irs

import "net/http"

type Status struct {
	HttpCode int
	Message  string
}

var ResponseStatuses = map[Code]Status{
	Ok: {http.StatusOK, "SUCCESS/OK"},

	BadRequest:                   {http.StatusBadRequest, "ERRORS/REQUEST/BAD_REQUEST"},
	Unauthorized:                 {http.StatusUnauthorized, "ERRORS/REQUEST/UNAUTHORIZED"},
	PaymentRequired:              {http.StatusPaymentRequired, "ERRORS/REQUEST/PAYMENT_REQUIRED"},
	Forbidden:                    {http.StatusForbidden, "ERRORS/REQUEST/FORBIDDEN"},
	NotFound:                     {http.StatusNotFound, "ERRORS/REQUEST/NOT_FOUND"},
	MethodNotAllowed:             {http.StatusMethodNotAllowed, "ERRORS/REQUEST/METHOD_NOT_ALLOWED"},
	NotAcceptable:                {http.StatusNotAcceptable, "ERRORS/REQUEST/NOT_ACCEPTABLE"},
	ProxyAuthRequired:            {http.StatusProxyAuthRequired, "ERRORS/REQUEST/PROXY_AUTH_REQUIRED"},
	RequestTimeout:               {http.StatusRequestTimeout, "ERRORS/REQUEST/REQUEST_TIMEOUT"},
	Conflict:                     {http.StatusConflict, "ERRORS/REQUEST/CONFLICT"},
	Gone:                         {http.StatusGone, "ERRORS/REQUEST/GONE"},
	LengthRequired:               {http.StatusLengthRequired, "ERRORS/REQUEST/LENGTH_REQUIRED"},
	PreconditionFailed:           {http.StatusPreconditionFailed, "ERRORS/REQUEST/PRECONDITION_FAILED"},
	RequestEntityTooLarge:        {http.StatusRequestEntityTooLarge, "ERRORS/REQUEST/REQUEST_ENTITY_TOO_LARGE"},
	RequestURITooLong:            {http.StatusRequestURITooLong, "ERRORS/REQUEST/REQUEST_URI_TOO_LONG"},
	UnsupportedMediaType:         {http.StatusUnsupportedMediaType, "ERRORS/REQUEST/UNSUPPORTED_MEDIA_TYPE"},
	RequestedRangeNotSatisfiable: {http.StatusRequestedRangeNotSatisfiable, "ERRORS/REQUEST/REQUESTED_RANGE_NOT_SATISFIABLE"},
	ExpectationFailed:            {http.StatusExpectationFailed, "ERRORS/REQUEST/EXPECTATION_FAILED"},
	Teapot:                       {http.StatusTeapot, "ERRORS/REQUEST/TEAPOT"},
	MisdirectedRequest:           {http.StatusMisdirectedRequest, "ERRORS/REQUEST/MISDIRECTED_REQUEST"},
	UnprocessableEntity:          {http.StatusUnprocessableEntity, "ERRORS/REQUEST/UNPROCESSABLE_ENTITY"},
	Locked:                       {http.StatusLocked, "ERRORS/REQUEST/LOCKED"},
	FailedDependency:             {http.StatusFailedDependency, "ERRORS/REQUEST/FAILED_DEPENDENCY"},
	TooEarly:                     {http.StatusTooEarly, "ERRORS/REQUEST/TOO_EARLY"},
	UpgradeRequired:              {http.StatusUpgradeRequired, "ERRORS/REQUEST/UPGRADE_REQUIRED"},
	PreconditionRequired:         {http.StatusPreconditionRequired, "ERRORS/REQUEST/PRECONDITION_REQUIRED"},
	TooManyRequests:              {http.StatusTooManyRequests, "ERRORS/REQUEST/TOO_MANY_REQUEST"},
	RequestHeaderFieldsTooLarge:  {http.StatusRequestHeaderFieldsTooLarge, "ERRORS/REQUEST/REQUEST_HEADER_FIELDS_TOO_LARGE"},
	UnavailableForLegalReasons:   {http.StatusUnavailableForLegalReasons, "ERRORS/REQUEST/UNAVAILABLE_FOR_LEGAL_REASONS"},

	InternalServerError:           {http.StatusInternalServerError, "ERRORS/REQUEST/INTERNAL_SERVER_ERROR"},
	NotImplemented:                {http.StatusNotImplemented, "ERRORS/REQUEST/NOT_IMPLEMENTED"},
	BadGateway:                    {http.StatusBadGateway, "ERRORS/REQUEST/BAD_GATEWAY"},
	ServiceUnavailable:            {http.StatusServiceUnavailable, "ERRORS/REQUEST/SERVICE_UNAVAILABLE"},
	GatewayTimeout:                {http.StatusGatewayTimeout, "ERRORS/REQUEST/GATEWAY_TIMEOUT"},
	HTTPVersionNotSupported:       {http.StatusHTTPVersionNotSupported, "ERRORS/REQUEST/HTTP_VERSION_NOT_SUPPORTED"},
	VariantAlsoNegotiates:         {http.StatusVariantAlsoNegotiates, "ERRORS/REQUEST/VARIANT_ALSO_NEGOTIATES"},
	InsufficientStorage:           {http.StatusInsufficientStorage, "ERRORS/REQUEST/INSUFFICIENT_STORAGE"},
	LoopDetected:                  {http.StatusLoopDetected, "ERRORS/REQUEST/LOOP_DETECTED"},
	NotExtended:                   {http.StatusNotExtended, "ERRORS/REQUEST/NOT_EXTENDED"},
	NetworkAuthenticationRequired: {http.StatusNetworkAuthenticationRequired, "ERRORS/REQUEST/NETWORK_AUTHENTICATION_REQUIRED"},
}

func AddStatus(code Code, httpCode int, message string) {
	ResponseStatuses[code] = Status{
		httpCode,
		message,
	}
}