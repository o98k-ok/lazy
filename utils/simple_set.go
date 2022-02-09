package utils

import "reflect"

type Set map[interface{}]bool

func AND(set1, set2 Set) Set {
	res := make(Set)

	for k, v := range set1 {
		if _, ok := set2[k]; !ok {
			continue
		}

		res[k] = v
	}
	return res
}

func OR(set1, set2 Set) Set {
	res := make(Set)
	for k, v := range set1 {
		res[k] = v
	}

	for k, v := range set2 {
		res[k] = v
	}
	return res
}

func ToSet(args interface{}) Set {
	res := make(Set)
	rv := reflect.ValueOf(args)
	if rv.Kind() != reflect.Slice {
		return res
	}

	for i := 0; i < rv.Len(); i++ {
		res[rv.Index(i).Interface()] = true
	}
	return res
}
