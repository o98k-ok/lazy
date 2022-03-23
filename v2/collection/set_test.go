package collection

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	arr := []string{"who", "i", "am"}
	brr := []string{"i", "am", "shadow"}

	set1 := NewSet[string](arr)
	set2 := NewSet[string](brr)

	t.Run("test and cases", func(t *testing.T) {
		res := set1.AND(set2)
		expect := NewSet([]string{"i", "am"})

		if !res.Equal(expect) {
			t.Errorf("test set failed, expect %v, got %v", expect, res)
		}
	})

	t.Run("test or cases", func(t *testing.T) {
		res := set1.OR(set2)
		expect := NewSet([]string{"i", "am", "shadow", "who"})

		if !res.Equal(expect) {
			t.Errorf("test set failed, expect %v, got %v", expect, res)
		}
	})

	t.Run("test for each cases", func(t *testing.T) {
		res := make([]string, 0)
		NewSet[string]([]string{"am", "i"}).ForEach(func(i string) {
			res = append(res, i)
		})

		if !reflect.DeepEqual(res, []string{"am", "i"}) && !reflect.DeepEqual(res, []string{"i", "am"}) {
			t.Errorf("test slice cases failed, got %v", res)
		}
	})
}
