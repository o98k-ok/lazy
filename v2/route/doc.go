package route

import (
	"github.com/brianvoe/gofakeit/v6"
)

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
	if err := gofakeit.Struct(&resp); err != nil {
		return "", err
	}

	elem := Elems{
		Method: method,
		URI:    path,
		Req:    req,
		Resp:   fn(resp),
	}
	return GenerateAPIDoc(elem)
}
