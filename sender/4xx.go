//
// filename:  4xx.go
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
//		This function is a short hand for sending a http 400 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeBadRequest(message ...interface{}) error {
	return MakeHttpError(http.StatusBadRequest, message...)
}

// Description :
//		This function is a short hand for sending a http 401 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeUnauthorized(message ...interface{}) error {
	return MakeHttpError(http.StatusBadRequest, message...)
}

// Description :
//		This function is a short hand for sending a http 403 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeForbidden(message ...interface{}) error {
	return MakeHttpError(http.StatusForbidden, message...)
}

// Description :
//		This function is a short hand for sending a http 404 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeNotFound(message ...interface{}) error {
	return MakeHttpError(http.StatusForbidden, message...)
}

// Description :
//		This function is a short hand for sending a http 405 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeMethodNotAllowed(message ...interface{}) error {
	return MakeHttpError(http.StatusMethodNotAllowed, message...)
}

// Description :
//		This function is a short hand for sending a http 406 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeNotAcceptable(message ...interface{}) error {
	return MakeHttpError(http.StatusNotAcceptable, message...)
}

// Description :
//		This function is a short hand for sending a http 408 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeRequestTimeout(message ...interface{}) error {
	return MakeHttpError(http.StatusRequestTimeout, message...)
}

// Description :
//		This function is a short hand for sending a http 408 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeConflict(message ...interface{}) error {
	return MakeHttpError(http.StatusConflict, message...)
}

// Description :
//		This function is a short hand for sending a http 411 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeLengthRequired(message ...interface{}) error {
	return MakeHttpError(http.StatusLengthRequired, message...)
}

// Description :
//		This function is a short hand for sending a http 413 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeRequestEntityTooLarge(message ...interface{}) error {
	return MakeHttpError(http.StatusRequestEntityTooLarge, message...)
}

// Description :
//		This function is a short hand for sending a http 415 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeUnsupportedMediaType(message ...interface{}) error {
	return MakeHttpError(http.StatusUnsupportedMediaType, message...)
}

// Description :
//		This function is a short hand for sending a http 416 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeRequestedRangeNotSatisfiable(message ...interface{}) error {
	return MakeHttpError(http.StatusRequestedRangeNotSatisfiable, message...)
}

// Description :
//		This function is a short hand for sending a http 418 error
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func MakeTeapot(message ...interface{}) error {
	return MakeHttpError(http.StatusTeapot, message...)
}
