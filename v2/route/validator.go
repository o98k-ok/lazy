package route

import (
	"github.com/go-playground/validator/v10"
)

var vlidtor = validator.New()

func Check[T any](v *T) error {
	return vlidtor.Struct(v)
}
