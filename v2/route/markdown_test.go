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

func TestNested(t *testing.T) {
	type Level2 struct {
		Age *string `json:"age" fake:"{word}"`
	}
	type Level1 struct {
		Name  string  `json:"name"  fake:"{firstname}"`
		Level *Level2 `json:"level"`
	}

	var l1, l2 Level1
	gofakeit.Struct(&l1)
	gofakeit.Struct(&l2)
	elems := Elems{
		Method: "GET",
		URI:    "/api/v1/users",
		Req:    l1,
		Resp:   l2,
	}

	fmt.Println(GenerateAPIDoc(elems))
}
