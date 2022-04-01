package format

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	cases := []struct {
		Name   string
		Input  interface{}
		Expect string
	}{
		{
			"test string slice",
			[]string{"shadowhao", "hyperma"},
			`[
    "shadowhao",
    "hyperma"
]`,
		},
		{
			"test int slice",
			[]int{10, 11, 12},
			`[
    10,
    11,
    12
]`,
		},
		{
			"test deep slice",
			[][]int{
				{10, 20, 30},
			},
			`[
    [
        10,
        20,
        30
    ]
]`,
		},
		{
			"test float slice",
			[]float64{10.0, 11.1, 12.2},
			`[
    10,
    11.1,
    12.2
]`,
		},
		{
			"test float slice",
			[]float32{10.0, 11.1, 12.2},
			`[
    10,
    11.1,
    12.2
]`,
		},
		{
			"test int",
			10,
			"10",
		},
		{
			"test null",
			nil,
			"null",
		},
		{
			"test simple map",
			map[string]int{
				"shadow": 10,
			},
			`{
    "shadow": 10
}`,
		},
		{
			"test hard map",
			map[string]map[string][]int{
				"shadow": {
					"math": []int{10, 20, 30},
				},
			},
			`{
    "shadow": {
        "math": [
            10,
            20,
            30
        ]
    }
}`,
		},
		{
			"test empty list",
			[]int{},
			`[]`,
		},
		{
			"test empty map",
			map[string]interface{}{},
			`{}`,
		},
		{
			"test complex case",
			map[string]interface{}{
				"shadow": nil,
			},
			`{
    "shadow": null
}`,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			decoder := NewEncoder(writer).DisableColor()
			err := decoder.Encode(c.Input)
			assert.NoError(t, err)

			got := writer.String()
			assert.Equal(t, c.Expect, got)
		})
	}

	t.Run("test empty indent cases", func(t *testing.T) {
		expect := "{\n->\"shadow\": \"ok\"\n}"
		writer := &bytes.Buffer{}
		decoder := NewEncoder(writer).DisableColor().WithIndent("->")
		err := decoder.Encode(map[string]string{"shadow": "ok"})
		assert.NoError(t, err)

		got := writer.String()
		assert.Equal(t, expect, got)
	})
}
