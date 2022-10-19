package collection

import (
	"fmt"
	"strings"
)

func SearchSlice[T any](items []T, key string) []T {
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

func SearchMap[T any](items map[string]T, key string) map[string]T {
	res := make(map[string]T)
	for key, item := range items {
		dest := strings.ToLower(fmt.Sprint(item))
		match := strings.ToLower(key)
		if strings.Contains(dest, match) {
			res[key] = item
		}
	}
	return res
}
