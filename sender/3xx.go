//
// filename:  3xx.go
// author:    Thomas Lombard
// copyright: Thomas Lombard
// license:   MIT
// status:    published
//
package irs

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Description :
//		This function is a short hand for sending a http 301 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		url string :
//			The url to redirect
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendMovedPermanently(c echo.Context, url string) error {
	return c.Redirect(http.StatusMovedPermanently, url)
}

// Description :
//		This function is a short hand for sending a http 302 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		url string :
//			The url to redirect
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendFound(c echo.Context, url string) error {
	return c.Redirect(http.StatusFound, url)
}

// Description :
//		This function is a short hand for sending a http 307 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		url string :
//			The url to redirect
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendTemporaryRedirect(c echo.Context, url string) error {
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// Description :
//		This function is a short hand for sending a http 308 response
// Parameters :
//      c echo.Context :
//          The current request context from echo. Must never be set to nil.
//		url string :
//			The url to redirect
// Returns :
// 		error :
// 			The encountered error if there has one. Nil otherwise
func SendPermanentRedirect(c echo.Context, url string) error {
	return c.Redirect(http.StatusPermanentRedirect, url)
}
