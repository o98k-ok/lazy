package collection

import "reflect"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](s []T) Set[T] {
	res := make(Set[T])

	for _, c := range s {
		res[c] = struct{}{}
	}
	return res
}

func (s Set[T]) AND(set Set[T]) Set[T] {
	res := make(Set[T])

	for k, v := range s {
		if _, ok := set[k]; !ok {
			continue
		}

		res[k] = v
	}
	return res
}

func (s Set[T]) ForEach(fn func(T)) {
	for key, _ := range s {
		fn(key)
	}
}

func (s Set[T]) OR(set Set[T]) Set[T] {
	res := make(Set[T])
	for k, v := range s {
		res[k] = v
	}

	for k, v := range set {
		res[k] = v
	}
	return res
}

func (s Set[T]) Equal(set Set[T]) bool {
	return reflect.DeepEqual(s, set)
}
