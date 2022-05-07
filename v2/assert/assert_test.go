package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	t.Run("test NoError->Equal->NotEqual", func(t *testing.T) {
		var err error
		r1, r2 := 10, 100

		err = IfNoError(err).ElIfEqual(10, r1).ElIfNotEqual(10, r2).Unwrap()
		if err != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})

	t.Run("test Equal->NotEqual", func(t *testing.T) {
		r1, r2 := 10, 100

		err := IfEqual(10, r1).ElIfEqual(10, r1).ElIfNotEqual(10, r2).Unwrap()
		if errors.Unwrap(err) != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})

	t.Run("test NotEqual->Equal", func(t *testing.T) {
		r1, r2 := 10, 100

		err := IfNotEqual(10, r2).ElIfEqual(10, r1).ElIfNotEqual(10, r2).Unwrap()
		if errors.Unwrap(err) != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})
}
