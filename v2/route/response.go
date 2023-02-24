package route

import (
	"github.com/kataras/iris/v12"
)

type ResponseEntity struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type ResponseHandlerImp struct {
	StatusMap map[error]int
	MsgMap    map[error]string

	DefaultMsg  string
	DefaultCode int
}

func (r *ResponseHandlerImp) OK(ctx iris.Context, data interface{}) {
	ctx.JSON(&ResponseEntity{
		Code:    r.StatusMap[nil],
		Data:    data,
		Message: r.MsgMap[nil],
	})
}

func (r *ResponseHandlerImp) ParamErr(ctx iris.Context, err error) {
	r.Failed(ctx, err)
}

func (r *ResponseHandlerImp) Failed(ctx iris.Context, err error) {
	var code int = r.DefaultCode
	var msg string = r.DefaultMsg
	if c, ok := r.StatusMap[err]; ok {
		code = c
	}
	if c, ok := r.MsgMap[err]; ok {
		msg = c
	}

	ctx.JSON(&ResponseEntity{
		Code:    code,
		Data:    nil,
		Message: msg + ":" + err.Error(),
	})
}

func (r *ResponseHandlerImp) ResponseEntity(en interface{}) interface{} {
	return ResponseEntity{
		Code:    0,
		Data:    en,
		Message: "",
	}
}
