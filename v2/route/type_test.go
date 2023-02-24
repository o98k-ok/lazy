package route

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("normal cases", func(t *testing.T) {
		type case1 struct {
			A string `json:"a"`
			B string `json:"b"`
			C string `schema:"c"`
		}
		tt := TypeConfig{
			NameTag: "json",
			DescTag: "desc",
		}

		tables := tt.GenerateTable(reflect.TypeOf(case1{}))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["case1"]) == 2
		})
		t.Log(tables)

		tt = TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		tables = tt.GenerateTable(reflect.TypeOf(case1{}))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["case1"]) == 1
		})
		t.Log(tables)
	})

	t.Run("array cases", func(t *testing.T) {
		type case1 struct {
			A string `json:"a"`
			B string `json:"b"`
			C string `schema:"c"`
		}
		tt := TypeConfig{
			NameTag: "json",
			DescTag: "desc",
		}

		tables := tt.GenerateTable(reflect.TypeOf([]*case1{}))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["[]case1"]) == 2
		})
		t.Log(tables)

		tt = TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		tables = tt.GenerateTable(reflect.TypeOf([]case1{}))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["[]case1"]) == 1
		})
		t.Log(tables)
	})

	t.Run("simple cases", func(t *testing.T) {
		tt := TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		var testabc string
		tables := tt.GenerateTable(reflect.TypeOf(testabc))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["string"]) == 0
		})

		t.Log(tables)
		tables = tt.GenerateTable(reflect.TypeOf(&testabc))
		assert.Condition(t, func() (success bool) {
			return len(tables) == 1 && len(tables["string"]) == 0
		})
		t.Log(tables)
	})
}
