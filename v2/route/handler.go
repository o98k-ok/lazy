package route

import "github.com/kataras/iris/v12"

// RequestParser parse request from iris
type RequestParser[T any] interface {
	Parse(ctx iris.Context, attrs map[string]interface{}) (T, error)
}

// ContextParser parse context from iris
type ContextParser[Context any] interface {
	Parse(ctx iris.Context) (Context, error)
}

// Handler handle http request
type Handler[T, Context, R any] func(ctx Context, req T) (R, error)

// ResponseHandler response http
type ResponseHandler interface {
	OK(ctx iris.Context, data interface{})
	ParamErr(ctx iris.Context, err error)
	Failed(ctx iris.Context, errs error)
	ResponseEntity(en interface{}) interface{}
}

// DocHandler doc api document
type DocHandler[T, R any] interface {
	DocIt(method string, path string, fn func(interface{}) interface{}) (string, error)
}

type DocStorer interface {
	Store(DocItem)
}
