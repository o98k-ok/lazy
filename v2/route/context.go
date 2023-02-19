package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/memstore"
)

type Context struct {
	Param map[string]memstore.Entry
}

type ContextParserImp struct{}

func (c ContextParserImp) Parse(ctx iris.Context) (Context, error) {
	var res Context = Context{Param: make(map[string]memstore.Entry)}
	for _, k := range ctx.Params().Store {
		res.Param[k.Key] = k
	}
	return res, nil
}
