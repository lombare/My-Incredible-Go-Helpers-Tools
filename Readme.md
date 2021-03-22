# MIGHT -> My Incredible Go Helpers Tools



## Purposes

This package is an assembly of tools I commonly uses. So feel free to contribute or give ideas of what to improve.


> - [MIGHT -> My Incredible Go Helpers Tools](#might----my-incredible-go-helpers-tools)
>    * [Purposes](#purposes)
>        + [Incredible env manager](#incredible-env-manager)
>            - [Other cool things](#other-cool-things)
>                * [Env loading](#env-loading)
>                * [Env manipulations](#env-manipulations)
>    * [Incredible response simplifier](#incredible-response-simplifier)
>        + [Subpackages](#subpackages)
>            - [Incredible Response Simplifier for Api : `irsa`](#incredible-response-simplifier-for-api----irsa-)
>                * [Wrapped Statuses](#wrapped-statuses)
>                * [Generic errors management](#generic-errors-management)
>                * [Custom errors management](#custom-errors-management)
>                    + [First the error is handled and is "normal" (Like a forbidden access)](#first-the-error-is-handled-and-is--normal---like-a-forbidden-access-)
>                    + [Secondly the error is handled but is critic (like a fail while requesting the database)](#secondly-the-error-is-handled-but-is-critic--like-a-fail-while-requesting-the-database-)
>                    + [Thirdly the error is not handled and is considered as critic](#thirdly-the-error-is-not-handled-and-is-considered-as-critic)


### Incredible env manager

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



## Incredible response simplifier

`irs`is a package that gives many functions that are intended to simplify the error management and allow an API to generalize every response.

By default `irs`responses are defined by this template :

```go
type responseTemplate struct {
   HttpCode   int           `json:"httpCode"`
   HttpStatus string        `json:"httpStatus"`
   Message    string        `json:"message"`
   Content    interface{}   `json:"content"`
}
```

- HttpCode is the code HTTP of the response
- HttpStatus is a stringified version of the HTTPCode (given by `http.StatusText()`)
- Message will contain a string that allow the client to get more informations about what has failed if there is a fail
- Content contain the response. Aka what the api respond



### Subpackages

####  Incredible Response Simplifier for Api : `irsa`

This package is build to respond JSON and XML resources (more incoming) as they're build for API.

##### Wrapped Statuses

So the most common Api responses are wrapped here is the list :

| Http Code             | Prototype                                                    |
| --------------------- | ------------------------------------------------------------ |
| 200 : Ok              | `func SendOk(c echo.Context, body interface{}, message ...interface{}) error` |
| 201 : Created         | `func SendCreated(c echo.Context, body interface{}, message ...interface{}) error` |
| 202 : Accepted        | `func SendAccepted(c echo.Context, body interface{}, message ...interface{}) error` |
| 204 : No Content      | `func SendNoContent(c echo.Context, body interface{}, message ...interface{}) error` |
| 205 : Reset Content   | `func SendResetContent(c echo.Context, body interface{}, message ...interface{}) error` |
| 206 : Partial Content | `func SendPartialContent(c echo.Context, body interface{}, message ...interface{}) error` |

| Http Code                | Prototype                                                    |
| ------------------------ | ------------------------------------------------------------ |
| 301 : Moved Permanently  | `func SendMovedPermanently(c echo.Context, url string) error` |
| 302 : Found              | `func SendFound(c echo.Context, url string) error`           |
| 307 : Temporary Redirect | `func SendTemporaryRedirect(c echo.Context, url string) error` |
| 308 : Permanent Redirect | `func SendPermanentRedirect(c echo.Context, url string) error` |

| Http Code                             | Prototype                                                    |
| ------------------------------------- | ------------------------------------------------------------ |
| 400 : Bad Request                     | `func SendBadRequest(c echo.Context, message ...interface{}) error` |
| 401 : Unauthorized                    | `func SendUnauthorized(c echo.Context, message ...interface{}) error` |
| 403 : Forbidden                       | `func SendForbidden(c echo.Context, message ...interface{}) error` |
| 404 : Not Found                       | `func SendNotFound(c echo.Context, message ...interface{}) error` |
| 405 : Method Not Allowed              | `func SendMethodNotAllowed(c echo.Context, message ...interface{}) error` |
| 406 : Not Acceptable                  | `func SendNotAcceptable(c echo.Context, message ...interface{}) error` |
| 408 : Request Timeout                 | `func SendRequestTimeout(c echo.Context, message ...interface{}) error` |
| 409 : Conflict                        | `func SendConflict(c echo.Context, message ...interface{}) error` |
| 411 : Length Required                 | `func SendLengthRequired(c echo.Context, message ...interface{}) error` |
| 413 : Request Entity Too Large        | `func SendRequestEntityTooLarge(c echo.Context, message ...interface{}) error` |
| 415 : Unsupported Media Type          | `func SendUnsupportedMediaType(c echo.Context, message ...interface{}) error` |
| 416 : Requested Range Not Satisfiable | `func SendRequestedRangeNotSatisfiable(c echo.Context, message ...interface{}) error` |
| 418 : Teapot                          | `func SendTeapot(c echo.Context, message ...interface{}) error` |

| Http Code                   | Prototype                                                    |
| --------------------------- | ------------------------------------------------------------ |
| 500 : Internal Server Error | `func SendInternalServerError(c echo.Context, message ...interface{}) error` |
| 501 : Not implemented       | `func SendNotImplemented(c echo.Context, message ...interface{}) error` |
| 503 : Service Unavailable   | `func SendServiceUnavailable(c echo.Context, message ...interface{}) error` |
| 507 : Insufficient Storage  | `func SendInsufficientStorage(c echo.Context, message ...interface{}) error` |

Note that it is possible to send any other type of error code via the wrapped method of theses calls :

`func Send(c echo.Context, status int, body interface{}, message ...interface{}) error`

This method will respond Json or XML depending on the `Accept`header.

You can explicitly send either Json or XML by calling theses methods :

- `func SendXML(c echo.Context, status int, payload interface{}, message ...interface{}) error`
- `func SendJSON(c echo.Context, status int, payload interface{}, message ...interface{}) error`



##### Generic errors management

It is also possible to skip all theses fastidious steps by calling `SendCode(c echo.Context, code int, payload ...interface{})`. This last method will take a look at the pre-registered responses and use them to respond.

> Note that the last parameter is a variadic argument only to make it optional. So if you place two values there only the first one will be take into account.

So let's take an example :

```go
// In our main.go
const (
	EverythingNice = iota + StatusPadding
    AnErrorOccurred
)

func main() {
	// Whatever ...
	irs.AddStatus(EverythingNice, http.StatusOk, "SUCCESS/OK").
    	AddStatus(AnErrorOccurred, http.StatusInternalError, "ERRORS/INTERNAL")
}

// In our controller 
func Controller() echo.HandlerFunc {
    return func(c echo.Context) error {
        
        if /* error */ {
            return irsa.SendCode(c, AnErrorOccurred)
        }
        
        return irsa.SendCode(c, EverythingNice, Data)
    }
}
```

In this case we shorten our error management to his strict minimal. And this method ensures that the errors codes will ever be the same.

You can also make the code differs but make them respond the same `httpCode`and `message`. That way the debug is simplified because it is possible to identify what was the error. Without giving any too sensible information.

##### Custom errors management

Because this is still not enough `irsa` and `irss` gives a method to directly sends errors.

This function can be used via 3 ways :

###### First the error is handled and is "normal" (Like a forbidden access)

```go
func Controller() {
// Stuff
return irsa.SendError(c, irs.MakeNormalError(http.StatusForbidden, "Forbidden access"))
}
```

That way a function called by a controller can setup a response error without have any access to the echo context. This allow error management from anywhere in the application.

###### Secondly the error is handled but is critic (like a fail while requesting the database)

```go
func Controller() {
// Stuff
return irsa.SendError(c, irs.MakeCriticalError(http.StatusInternalError, err, "An error occurred"))
}
```

That makes a response with an http 500 code. But this will add `X-Error-Id`header that identifies a log in the api. That way any developer can have access to the full log of the error without any risk of exposing any sensible information.

###### Thirdly the error is not handled and is considered as critic

```go
func Controller() {
// Stuff
return irsa.SendError(c, err)
}
```

In this case the http Code is systematically 500 and the message is `Internal error` (because the error wasn't properly handled). And exactly like any critical error an `X-Error-Id`is issued and the full error is logged.