package irsa

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Description :
//		This function is a short hand for sending a http 400 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendBadRequest(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusBadRequest, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 401 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendUnauthorized(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusBadRequest, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 403 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendForbidden(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusForbidden, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 404 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendNotFound(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusForbidden, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 405 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendMethodNotAllowed(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusMethodNotAllowed, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 406 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendNotAcceptable(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusNotAcceptable, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 408 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendRequestTimeout(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusRequestTimeout, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 408 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendConflict(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusConflict, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 411 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendLengthRequired(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusLengthRequired, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 413 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendRequestEntityTooLarge(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusRequestEntityTooLarge, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 415 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendUnsupportedMediaType(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusUnsupportedMediaType, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 416 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendRequestedRangeNotSatisfiable(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusRequestedRangeNotSatisfiable, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 418 response
// Parameters :
//      c echo.Context : Required
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} : Optional
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendTeapot(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusTeapot, nil, message...)
}
