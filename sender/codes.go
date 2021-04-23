//
// filename:  codes.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irs

import (
	"fmt"
)

type Code int

func (c Code) Error() string {
	return fmt.Sprint("httpCode:", ResponseStatuses[c].HttpCode, ". message:", ResponseStatuses[c].Message)
}

func (c Code) HttpCode() int {
	v, ok := ResponseStatuses[c]
	if ok {
		return v.HttpCode
	}
	return -1
}

func (c Code) Message() string {
	v, ok := ResponseStatuses[c]
	if ok {
		return v.Message
	}
	return "ERRORS/UNKNOWN"
}

const (
	Ok Code = iota

	// Theses codes are defaults errcode for http codes.
	// This allow IRSA to handle echo errors and provide nice codes
	//
	// Some (iota + weird numbers) are required because every 400 codes does not follows
	// (So we step over them)
	BadRequest                   Code = iota + 399 // RFC 7231, 6.5.1
	Unauthorized                                   // RFC 7235, 3.1
	PaymentRequired                                // RFC 7231, 6.5.2
	Forbidden                                      // RFC 7231, 6.5.3
	NotFound                                       // RFC 7231, 6.5.4
	MethodNotAllowed                               // RFC 7231, 6.5.5
	NotAcceptable                                  // RFC 7231, 6.5.6
	ProxyAuthRequired                              // RFC 7235, 3.2
	RequestTimeout                                 // RFC 7231, 6.5.7
	Conflict                                       // RFC 7231, 6.5.8
	Gone                                           // RFC 7231, 6.5.9
	LengthRequired                                 // RFC 7231, 6.5.10
	PreconditionFailed                             // RFC 7232, 4.2
	RequestEntityTooLarge                          // RFC 7231, 6.5.11
	RequestURITooLong                              // RFC 7231, 6.5.12
	UnsupportedMediaType                           // RFC 7231, 6.5.13
	RequestedRangeNotSatisfiable                   // RFC 7233, 4.4
	ExpectationFailed                              // RFC 7231, 6.5.14
	Teapot                                         // RFC 7168, 2.3.3
	MisdirectedRequest           Code = iota + 401 // RFC 7540, 9.1.2
	UnprocessableEntity                            // RFC 4918, 11.2
	Locked                                         // RFC 4918, 11.3
	FailedDependency                               // RFC 4918, 11.4
	TooEarly                                       // RFC 8470, 5.2.
	UpgradeRequired                                // RFC 7231, 6.5.15
	PreconditionRequired         Code = iota + 402 // RFC 6585, 3
	TooManyRequests                                // RFC 6585, 4
	RequestHeaderFieldsTooLarge  Code = iota + 403 // RFC 6585, 5
	UnavailableForLegalReasons   Code = iota + 422 // RFC 7725, 3

	InternalServerError           Code = iota + 470 // RFC 7231, 6.6.1
	NotImplemented                                  // RFC 7231, 6.6.2
	BadGateway                                      // RFC 7231, 6.6.3
	ServiceUnavailable                              // RFC 7231, 6.6.4
	GatewayTimeout                                  // RFC 7231, 6.6.5
	HTTPVersionNotSupported                         // RFC 7231, 6.6.6
	VariantAlsoNegotiates                           // RFC 2295, 8.1
	InsufficientStorage                             // RFC 4918, 11.5
	LoopDetected                                    // RFC 5842, 7.2
	NotExtended                   Code = iota + 471 // RFC 2774, 7
	NetworkAuthenticationRequired                   // RFC 6585, 6

	StatusPadding = 1000
)
