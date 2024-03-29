package route

import (
	"reflect"

	"github.com/brianvoe/gofakeit/v6"
)

type DocItem struct {
	URI   string
	Group string
	Doc   string
	Tags  []string
}

type DocHandlerImp[T, R any] struct{}

func NewDocAPI[T, R any]() DocHandler[T, R] {
	return &DocHandlerImp[T, R]{}
}

func (d *DocHandlerImp[T, R]) DocIt(method string, path string, fn func(interface{}) interface{}) (string, error) {
	var req T
	var resp R
	if err := gofakeit.Struct(&req); err != nil {
		return "", err
	}

	if reflect.TypeOf(resp).Kind() == reflect.Slice {
		gofakeit.Slice(&resp)
	} else {
		if err := gofakeit.Struct(&resp); err != nil {
			return "", err
		}
	}

	elem := Elems{
		Method: method,
		URI:    path,
		Req:    req,
		Resp:   resp,
		Fn:     fn,
	}
	return GenerateAPIDoc(elem)
}
