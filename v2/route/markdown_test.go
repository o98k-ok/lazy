package route

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestMarkdown(t *testing.T) {
	var req entity
	var resp entity
	gofakeit.Struct(&req)
	gofakeit.Struct(&resp)
	elems := Elems{
		Method: "GET",
		URI:    "/api/v1/users",
		Req:    req,
		Resp:   resp,
	}

	fmt.Println(GenerateAPIDoc(elems))
}
