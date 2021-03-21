package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Description :
//		This function is a short hand for sending a http 500 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendInternalServerError(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusInternalServerError, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 501 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendNotImplemented(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusNotImplemented, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 503 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendServiceUnavailable(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusServiceUnavailable, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 507 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendInsufficientStorage(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusInsufficientStorage, nil, message...)
}
