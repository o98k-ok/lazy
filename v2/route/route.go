package route

import (
	"fmt"
	"path/filepath"

	"github.com/kataras/iris/v12"
)

type Group[Context any] struct {
	Context ContextParser[Context]
	Resp    ResponseHandler
	App     iris.Party
}

type Handlers[T, C, R any] struct {
	Request RequestParser[T]
	API     DocHandler[T, R]
	Handler Handler[T, C, R]
}

func NewHandler[T, C, R any](handler Handler[T, C, R]) *Handlers[T, C, R] {
	return &Handlers[T, C, R]{
		Request: NewRequest[T](),
		API:     NewDocAPI[T, R](),
		Handler: handler,
	}
}

func Route[T, C, R any](group Group[C], method string, relativePath string, hs *Handlers[T, C, R], attrs map[string]interface{}) {
	real := func(ctx iris.Context) {
		// get context
		var context C
		var err error
		if group.Context != nil {
			context, err = group.Context.Parse(ctx)
			if err != nil {
				group.Resp.ParamErr(ctx, err)
				return
			}
		}

		// get request
		req, err := hs.Request.Parse(ctx, attrs)
		if err != nil {
			group.Resp.ParamErr(ctx, err)
			return
		}

		// main handler and
		resp, err := hs.Handler(context, req)
		if err == nil {
			group.Resp.OK(ctx, resp)
			return
		}
		group.Resp.Failed(ctx, err)
	}

	path := group.App.GetRelPath()
	fmt.Println(hs.API.DocIt(method, filepath.Join(path, relativePath)))
	group.App.Handle(method, relativePath, real)
}
