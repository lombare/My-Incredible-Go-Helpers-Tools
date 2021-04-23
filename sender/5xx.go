//
// filename:  5xx.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irs

import (
	"net/http"
)

// Description :
//		This function is a short hand for sending a http 500 error
// Parameters :
//      err error : Required
//          The error that causes the request fail. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeInternalServerError(err error, message ...interface{}) error {
	return MakeServerError(http.StatusInternalServerError, err, message...)
}

// Description :
//		This function is a short hand for sending a http 501 error
// Parameters :
//      err error : Required
//          The error that causes the request fail. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeNotImplemented(err error, message ...interface{}) error {
	return MakeServerError(http.StatusNotImplemented, err, message...)
}

// Description :
//		This function is a short hand for sending a http 503 error
// Parameters :
//      err error : Required
//          The error that causes the request fail. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeServiceUnavailable(err error, message ...interface{}) error {
	return MakeServerError(http.StatusServiceUnavailable, err, message...)
}

// Description :
//		This function is a short hand for sending a http 507 error
// Parameters :
//      err error : Required
//          The error that causes the request fail. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeInsufficientStorage(err error, message ...interface{}) error {
	return MakeServerError(http.StatusInsufficientStorage, err, message...)
}
