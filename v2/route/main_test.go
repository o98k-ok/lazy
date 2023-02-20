package route

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/kataras/iris/v12"
)

type entity struct {
	Name *string `json:"name" schema:"name" desc:"姓名" fake:"{firstname}"`
}

func Hello(ctx Context, req entity) (entity, error) {
	id, _ := ctx.Param["id"].Uint32Default(0)
	name := fmt.Sprintf("id: %d; name: %s", id, *req.Name)
	return entity{
		Name: &name,
	}, nil
}

func TestLogic(t *testing.T) {
	app := iris.New()
	group := Group[Context]{
		Context: ContextParserImp{},
		Resp:    &ResponseHandlerImp{DefaultMsg: "ok", DefaultCode: 0},
		App:     app,
	}
	var val string = "o98k"
	Route(group, http.MethodGet, "/hello/{id:uint32}", NewHandler(Hello), map[string]interface{}{DEFAULT: entity{Name: &val}})
	// app.Listen(":8080")
}
