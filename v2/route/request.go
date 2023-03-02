package route

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/o98k-ok/schema"
)

type Request interface {
	ForceParseFromParam() bool
	ValidRequest() bool
	InitRequest()
}

const (
	// AUTH  = "with_auth"
	PARAM   = "froce_param"   // orce parse request from http user param. default judge by method
	VALID   = "valid_request" // valid request body or request param via schema
	DEFAULT = "default_value" // parse request with default value
)

func (r *RequestParserImp[T]) ForceParseFromParam() bool {
	return r.Attrs[PARAM].(bool)
}

func (r *RequestParserImp[T]) ValidRequest() bool {
	return r.Attrs[VALID].(bool)
}

func NewRequest[T any]() RequestParser[T] {
	return &RequestParserImp[T]{
		Attrs: map[string]interface{}{
			PARAM: false,
			VALID: true,
		},
	}
}

type RequestParserImp[T any] struct {
	Attrs map[string]interface{}
}

func (rp *RequestParserImp[T]) Parse(ctx iris.Context, attrs map[string]interface{}) (T, error) {
	// set request attrs
	for k, v := range attrs {
		rp.Attrs[k] = v
	}

	var param bool
	if rp.ForceParseFromParam() || ctx.Method() == http.MethodGet {
		param = true
	}

	var req T
	if param {
		decoder := schema.NewDecoder()
		err := decoder.Decode(req, ctx.Request().URL.Query())
		if err != nil {
			return req, err
		}
	} else {
		if err := ctx.ReadJSON(&req); err != nil {
			return req, err
		}
	}

	if rp.ValidRequest() {
		if err := Check(&req); err != nil {
			return req, err
		}
	}
	return req, nil
}
