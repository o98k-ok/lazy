package utils

import "reflect"

type Set map[interface{}]bool

func NewSet(s interface{}) Set {
	res := make(Set)

	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Slice {
		return nil
	}

	for i := 0; i < val.Len(); i++ {
		res[val.Index(i).Interface()] = true
	}
	return res
}

func (s Set) AND(set Set) Set {
	res := make(Set)

	for k, v := range s {
		if _, ok := set[k]; !ok {
			continue
		}

		res[k] = v
	}
	return res
}

func (s Set) ForEach(fn func(interface{})) {
	for key, _ := range s {
		fn(key)
	}
}

func (s Set) OR(set Set) Set {
	res := make(Set)
	for k, v := range s {
		res[k] = v
	}

	for k, v := range set {
		res[k] = v
	}
	return res
}

func (s Set) Equal(set Set) bool {
	return reflect.DeepEqual(s, set)
}
