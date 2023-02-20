package route

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/kataras/iris/v12"
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

func (r *RequestParserImp[T]) InitRequest() {
	val, ok := r.Attrs[DEFAULT]
	if !ok {
		return
	}

	v, ok := val.(T)
	if !ok {
		return
	}
	r.Req = v
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
	Req   T
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

	rp.InitRequest()
	if param {
		decoder := schema.NewDecoder()
		err := decoder.Decode(&rp.Req, ctx.Request().URL.Query())
		if err != nil {
			return rp.Req, err
		}
	} else {
		if err := ctx.ReadJSON(&rp.Req); err != nil {
			return rp.Req, err
		}
	}

	if rp.ValidRequest() {
		if err := Check(&rp.Req); err != nil {
			return rp.Req, err
		}
	}
	return rp.Req, nil
}
