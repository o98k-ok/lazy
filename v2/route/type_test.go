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

		table, name := tt.GenerateTable(reflect.TypeOf(case1{}))
		assert.Condition(t, func() (success bool) {
			return len(table) == 2 && name == "case1"
		})

		tt = TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		table, name = tt.GenerateTable(reflect.TypeOf(case1{}))
		assert.Condition(t, func() (success bool) {
			return len(table) == 1 && name == "case1"
		})
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

		table, name := tt.GenerateTable(reflect.TypeOf([]*case1{}))
		assert.Condition(t, func() (success bool) {
			return len(table) == 2 && name == "[]case1"
		})

		tt = TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		table, name = tt.GenerateTable(reflect.TypeOf([]case1{}))
		assert.Condition(t, func() (success bool) {
			return len(table) == 1 && name == "[]case1"
		})
	})

	t.Run("simple cases", func(t *testing.T) {
		tt := TypeConfig{
			NameTag: "schema",
			DescTag: "desc",
		}
		var testabc string
		table, name := tt.GenerateTable(reflect.TypeOf(testabc))
		assert.Condition(t, func() (success bool) {
			return len(table) == 0 && name == "string"
		})
		table, name = tt.GenerateTable(reflect.TypeOf(&testabc))
		assert.Condition(t, func() (success bool) {
			return len(table) == 0 && name == "string"
		})
	})
}
