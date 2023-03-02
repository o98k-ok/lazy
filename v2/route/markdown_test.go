package route

import (
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/o98k-ok/schema"
	"github.com/stretchr/testify/assert"
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

func TestDefalultForSchema(t *testing.T) {
	type Abc struct {
		Int int `schema:"int" default:"10"`
	}

	var a Abc
	decoder := schema.NewDecoder()
	err := decoder.Decode(&a, map[string][]string{})
	assert.NoError(t, err)
	assert.Equal(t, a.Int, 10)

	err = decoder.Decode(&a, map[string][]string{"int": {"20"}})
	assert.NoError(t, err)
	assert.Equal(t, a.Int, 20)
}
