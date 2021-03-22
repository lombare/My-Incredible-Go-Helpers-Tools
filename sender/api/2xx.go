package irsa

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Description :
//		This function is a short hand for sending a http 200 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		body interface{} :
//			The body to send as response
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendOk(c echo.Context, body interface{}, message ...interface{}) error {
	return Send(c, http.StatusOK, body, message...)
}

// Description :
//		This function is a short hand for sending a http 201 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		body interface{} :
//			Must be the created ressource.
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendCreated(c echo.Context, body interface{}, message ...interface{}) error {
	return Send(c, http.StatusCreated, body, message...)
}

// Description :
//		This function is a short hand for sending a http 202 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		body interface{} :
//			The body to send as response if there must be one.
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendAccepted(c echo.Context, body interface{}, message ...interface{}) error {
	return Send(c, http.StatusAccepted, body, message...)
}

// Description :
//		This function is a short hand for sending a http 204 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendNoContent(c echo.Context, message ...interface{}) error {
	return Send(c, http.StatusNoContent, nil, message...)
}

// Description :
//		This function is a short hand for sending a http 205 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		body interface{} :
//			The new ressource value
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendResetContent(c echo.Context, body interface{}, message ...interface{}) error {
	return Send(c, http.StatusResetContent, body, message...)
}

// Description :
//		This function is a short hand for sending a http 206 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		body interface{} :
//			The part of the response
//		message ...interface{} :
//			Informations about the response. Used to fill the message field of the response template.
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendPartialContent(c echo.Context, body interface{}, message ...interface{}) error {
	return Send(c, http.StatusPartialContent, body, message...)
}
