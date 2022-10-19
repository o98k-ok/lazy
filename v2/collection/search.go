package collection

import (
	"fmt"
	"strings"
)

func Search[T any](items []T, key string) []T {
	var res []T
	for _, item := range items {
		dest := strings.ToLower(fmt.Sprint(item))
		match := strings.ToLower(key)
		if strings.Contains(dest, match) {
			res = append(res, item)
		}
	}
	return res
}
