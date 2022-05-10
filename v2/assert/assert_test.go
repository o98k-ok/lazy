package assert

import (
	"errors"
	"testing"
)

func TestAssert(t *testing.T) {
	t.Run("test NoError->Equal->NotEqual", func(t *testing.T) {
		var err error
		r1, r2 := 10, 100

		err = NoError(err).AndEqual(10, r1).AndNotEqual(10, r2).Unwrap()
		if err != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})

	t.Run("test Equal->NotEqual", func(t *testing.T) {
		r1, r2 := 10, 100

		err := Equal(10, r1).AndEqual(10, r1).AndNotEqual(10, r2).Unwrap()
		if errors.Unwrap(err) != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})

	t.Run("test NotEqual->Equal", func(t *testing.T) {
		r1, r2 := 10, 100

		err := NotEqual(10, r2).AndEqual(10, r1).AndNotEqual(10, r2).Unwrap()
		if errors.Unwrap(err) != nil {
			t.Fatalf("test failed, err %v", err)
		}
	})
}
