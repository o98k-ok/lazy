package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAttr(t *testing.T) {
	t.Run("normal cases", func(t *testing.T) {
		d := map[string]interface{}{
			"name": "shadow",
			"inner": map[string]interface{}{
				"age": 10,
			},
		}

		res, err := GetAttr([]string{"name"}, d)
		assert.NoError(t, err)

		val, ok := res.(string)
		assert.True(t, ok)
		assert.Equal(t, "shadow", val)

		res, err = GetAttr([]string{"inner", "age"}, d)
		assert.NoError(t, err)

		val2, ok2 := res.(int)
		assert.True(t, ok2)
		assert.Equal(t, 10, val2)

		res, err = GetAttr([]string{"inner", "age"}, nil)
		assert.NoError(t, err)
		assert.Nil(t, res)
	})
}
