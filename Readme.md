# MIGHT -> My Incredible Go Helpers Tools

> - [MIGHT -> My Incredible Go Helpers Tools](#might----my-incredible-go-helpers-tools)
>    * [Purposes](#purposes)
>    * [Still To Do](#still-to-do)
>    * [Incredible env manager `iem`](#incredible-env-manager)
>        + [Other cool things](#other-cool-things)
>            - [Env loading](#env-loading)
>            - [Env manipulations](#env-manipulations)
>    * [Incredible response simplifier `irs`](#incredible-response-simplifier-irs)
>        + [Generic errors management](#generic-errors-management)
>            - [Create a generic error](#create-a-generic-error)
>            - [Send a generic error](#send-a-generic-error)
>            - [Wrapped methods of `irs`](#wrapped-methods-of-irs)
>        + [Incredible Response Simplifier for Api : `irsa`](#incredible-response-simplifier-for-api--irsa)
>        + [Incredible Response Simplifier for Static : `irss`](#incredible-response-simplifier-for-static--irss)
>        + [Wrapped Statuses by both `irsa` and `irss`](#wrapped-statuses-by-both-irsa-and-irss)
>            - [Using generic errors](#using-generic-errors)
>            - [Custom errors management](#custom-errors-management)
>                    + [MakeHttpError](#makehttperror)
>                    + [MakeServerError](#makeservererror)
>                    + [MakeCodeError](#makecodeerror)
>                    + [Unhandled error](#unhandled-error)
## Purposes

This package is an assembly of tools I commonly uses. So feel free to contribute.

## Still to do

- [x] Load env from JSON
- [ ] Load env from XML
- [x] Load env from Dotenv `.env`
- [ ] Load env from Yaml
- [x] Clear environment
- [x] Native type getters
- [ ] Config fallbacks
- [ ] Variable that can only be retrieved from host environment
- [ ] Variable that can only be retrieved from config file 
- [ ] Variable that can only be retrieved from command line
- [ ] Check env config
- [ ] Load env to a struct
- [ ] Handle nested vars (with struct support)

## Incredible env manager `iem`

`iem`allows you to manage your programs configurations via the simplest way. The package allows you to get any value via many type of getters.

Each getter is available in 3 versions

- `Get___(varname string)` : This method returns the variable, with an error if the value does not exist or if the conversion fails.
- `GetDefault___(varname string, def any)`: This type of getter returns the variable if it exist or the default value otherwise.
- `MustGet___(varname string)` : This type of getter returns the value if the variable exists. Panic otherwise.

This 3 getters are available in 18 versions :

- `bool`
- `int` & `uint`
- `int8` & `uint8`
- `int16` & `uint16`
- `int32` & `uint32`
- `int64` & `uint64`
- `float32`& `float64`
- `complex64`& `complex128`
- `string` (Just called `Get`, `GetDefault` and `MustGet`)
- `file` Theses handlers returns `io.ReadWriteCloser`
- `fileContent` Theses getters returns a `[]byte` containing the file content.



#### Other cool things

##### Env loading

You can load a file via 2 syntaxes (more incoming). Dotenv and JSON.

Once again you can load a file via two methods : `Load(filename string)`, `MustLoad(filename string)`.

For three kinds of loads :

- `Load(filename string)` & `MustLoad(filename string)`: load the file in `.env`or in `.env.json`format.
- `LoadDotenv(filename string)`&`MustLoadDotenv(filename string)`: load `.env`files
- `LoadJsonEnv(filename string)`&`MustLoadJsonEnv(filename string)`: load `.env.json`files.



##### Env manipulations

Theses functions adds nothing more as they just wrap some golang native calls. But they're here to make everything uniform.

- `Has(varname string)`: This function returns a boolean that tells if a variable is set in the environment.
- `Set(varname, value string)`: This function sets a variable in the env.
- `ClearEnv()`: This function clear the environment



## Incredible response simplifier `irs`

`irs` is a package that gives many functions that are intended to simplify the error management and allow an API to generalize every response.

By default `irs` responses are defined by this template when there is no error:

```go
type responseTemplate struct {
    Content    interface{} `json:"content"`
}
```
- Content contain the response. Aka what the api respond

```go
type errorResponseTemplate struct {
    ErrCode    irs.Code `json:"errCode"`
    Message    string   `json:"message"`
}
```

- ErrCode is the error code of the response (the one in the enum)
- Message will contain a string that allow the client to get more information about what has failed if there is a fail


### Generic errors management

`irs` allows you to make some generic errors that you can call from anywhere in your code. That way you are sure that 
the errors you send will ever be the same, with the same message, and the same code.

#### Create a generic error 

Do just like this :

```go
package whatever

import (
	"net/http"

	irs "github.com/lombare/might/sender"
)

const (
	// It is important to pad your messages. That way you'll never collide any other error type
	AnyErrorCode irs.Code = iota + irs.StatusPadding
)

func Whatever() {
	irs.AddStatus(AnyErrorCode, http.StatusBadRequest, "You send a bad request because ...")
}
```

#### Send a generic error 

Just returns the code error :

```go
// In our controller 
func Controller() echo.HandlerFunc {
    return func(c echo.Context) error {
        return AnyErrorCode
    }
}
```

#### Wrapped methods of `irs`

| Http Code                | Prototype                                                    |
| ------------------------ | ------------------------------------------------------------ |
| 301 : Moved Permanently  | `func SendMovedPermanently(c echo.Context, url string) error` |
| 302 : Found              | `func SendFound(c echo.Context, url string) error`           |
| 307 : Temporary Redirect | `func SendTemporaryRedirect(c echo.Context, url string) error` |
| 308 : Permanent Redirect | `func SendPermanentRedirect(c echo.Context, url string) error` |

| Http Code                             | Prototype                                                    |
| ------------------------------------- | ------------------------------------------------------------ |
| 400 : Bad Request                     | `func MakeBadRequest(message ...interface{}) error` |
| 401 : Unauthorized                    | `func MakeUnauthorized(message ...interface{}) error` |
| 403 : Forbidden                       | `func MakeForbidden(message ...interface{}) error` |
| 404 : Not Found                       | `func MakeNotFound(message ...interface{}) error` |
| 405 : Method Not Allowed              | `func MakeMethodNotAllowed(message ...interface{}) error` |
| 406 : Not Acceptable                  | `func MakeNotAcceptable(message ...interface{}) error` |
| 408 : Request Timeout                 | `func MakeRequestTimeout(message ...interface{}) error` |
| 409 : Conflict                        | `func MakeConflict(message ...interface{}) error` |
| 411 : Length Required                 | `func MakeLengthRequired(message ...interface{}) error` |
| 413 : Request Entity Too Large        | `func MakeRequestEntityTooLarge(message ...interface{}) error` |
| 415 : Unsupported Media Type          | `func MakeUnsupportedMediaType(message ...interface{}) error` |
| 416 : Requested Range Not Satisfiable | `func MakeRequestedRangeNotSatisfiable(message ...interface{}) error` |
| 418 : Teapot                          | `func MakeTeapot(message ...interface{}) error` |

| Http Code                   | Prototype                                                    |
| --------------------------- | ------------------------------------------------------------ |
| 500 : Internal Server Error | `func MakeInternalServerError(message ...interface{}) error` |
| 501 : Not implemented       | `func MakeNotImplemented(message ...interface{}) error` |
| 503 : Service Unavailable   | `func MakeServiceUnavailable(message ...interface{}) error` |
| 507 : Insufficient Storage  | `func MakeInsufficientStorage(message ...interface{}) error` |


## Incredible Response Simplifier for Api : `irsa`

This package is build to respond JSON and XML resources (more incoming) as they're build for API.

## Incredible Response Simplifier for Static : `irss`

This package is build to respond static files resources as it is build to send static files

### Wrapped Statuses by both `irsa` and `irss`

So the most common Api responses are wrapped here is the list :

| Http Code             | Prototype                                                    |
| --------------------- | ------------------------------------------------------------ |
| 200 : Ok              | `func SendOk(c echo.Context, body interface{}, message ...interface{}) error` |
| 201 : Created         | `func SendCreated(c echo.Context, body interface{}, message ...interface{}) error` |
| 202 : Accepted        | `func SendAccepted(c echo.Context, body interface{}, message ...interface{}) error` |
| 204 : No Content      | `func SendNoContent(c echo.Context, body interface{}, message ...interface{}) error` |
| 205 : Reset Content   | `func SendResetContent(c echo.Context, body interface{}, message ...interface{}) error` |
| 206 : Partial Content | `func SendPartialContent(c echo.Context, body interface{}, message ...interface{}) error` |

Note that it is possible to send any other type of http code via the method that these functions wrap :

```go
    func Send(c echo.Context, status int, body interface{}) error
```

This method will respond Json or XML depending on the `Accept` header (More incoming).

##### Using generic errors

It is also possible to skip all theses fastidious steps by calling `SendCode(c echo.Context, code int, payload ...interface{})`. This last method will take a look at the pre-registered responses and use them to respond.

> Note that the last parameter is a variadic argument only to make it optional. So if you place two values there only the first one will be take into account.

So let's take an example :

```go
// In our main.go
const (
    AnErrorOccurred irs.Code = iota + StatusPadding
)

func main() {
	// Whatever ...
    AddStatus(AnErrorOccurred, http.StatusInternalError, "ERRORS/INTERNAL")
}

// In our controller 
func Controller() echo.HandlerFunc {
    return func(c echo.Context) error {
        
        if /* error */ {
            return AnErrorOccurred
        }
        
        return irsa.SendOk(c, Data)
    }
}
```

In this case we shorten our error management to his strict minimal. And this method ensures that the errors codes will ever be the same.

You can also make the `errCode` differs but send the same `httpCode`and `message`. That way the debug is simplified because it is possible to identify what was the error without giving any too sensible information leak.

##### Custom errors management

Because this is still not enough `irs` gives a method to make different kinds of errors.

There is 3 functions that allows you to make errors from anywhere.

###### MakeHttpError

`MakeHttpError` is a function that makes an error that is intended to explain a normal http error. A normal error can be a forbiden access
for a user that is signed in but has no rights to access something. This is an error that doesn't need any particular attention.

Here is an example :
```go

func doSomeVerifications() {
    // Stuff here 
    if !ok {
        irs.MakeHttpError(http.StatusForbidden, "Forbidden access")
    }
    reutrn nil
}

func Controller(c echo.context) {

    err := doSomeVerifications()
    if err != nil {
        return err
    } 
    return irsa.SendOk(c, "whatever")
}
```

The function `doSomeVerifications()` does not have any access to any echo stuff but is now able to throw a specific error.

###### MakeServerError

`MakeServerError` is a function that makes an error that requires developer attention. A good example for this kind of errors can be a database access failure.

Here is an example of application

```go
package whatever

import (
	"net/http"

	"github.com/labstack/echo/v4"
	irs "github.com/lombare/might/sender"
	irsa "github.com/lombare/might/sender/api"
)

func doSomeVerifications() error {
	// Stuff here
	if !ok {
		return irs.MakeHttpError(http.StatusForbidden, "ERRORS/REQUEST/FORBIDDEN_ACCESS")
	}
	return nil
}

func getSomethingInDatabase() (interface{}, error) {

	// Suff here
	data, err := databaseCallThatCanFail()
	if err != nil {
		return nil, irs.MakeInternalServerError(err, "ERRORS/SERVER/INTERNAL")
	}

	return data, nil
}

func Controller(c echo.Context) error {

	err := doSomeVerifications()
	if err != nil {
		return err
	}

	data, err := getSomethingInDatabase()
	if err != nil {
		return err
	}

	return irsa.SendOk(c, data)
}
```

That makes a response with an http 500 code. But this will add `X-Error-Id` header that identifies a log in the api. 
That way any developer can have access to the full log of the error without any risk of exposing any sensible information in production.

###### MakeCodeError

This function is made for the really common errors. Such as a not found or anything else.
The errors building is reduces to his strict minimum :

```go
package whatever

import (
	"net/http"

	"github.com/labstack/echo/v4"
	irs "github.com/lombare/might/sender"
	irsa "github.com/lombare/might/sender/api"
)


func doSomeVerifications() error {
	// Stuff here
	if !ok {
		return irs.MakeForbidden()
	}
	return nil
}

func getSomethingInDatabase() (interface{}, error) {

	// Suff here
	data, err := databaseCallThatCanFails()
	if err != nil {
		return nil, irs.MakeInternalServerError(err)
	}

	// Here we shortened the error to his simplest expression
	if /* Database failure */ {
		return nil, ReallyCommonDatabaseError
	}

	return data, nil
}

func Controller(c echo.Context) error {

	err := doSomeVerifications()
	if err != nil {
		return err
	}

	data, err := getSomethingInDatabase()
	if err != nil {
		return err
	}

	return irsa.SendOk(c, data)
}
```

###### Unhandled error

In this case the error will be considered as a server error. The HttpCode will ever be 500 and the message will ever be `ERRORS/SERVER/INTERNAL`
