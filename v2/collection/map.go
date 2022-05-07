package collection

import (
	"fmt"
)

func GetAttr(keys []string, attr interface{}) (interface{}, error) {
	if attr == nil {
		return nil, nil
	}

	for _, c := range keys {
		val, ok := attr.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("no such keys %v", keys)
		}

		attr, ok = val[c]
		if !ok {
			return nil, fmt.Errorf("no such keys %v", keys)
		}
	}
	return attr, nil
}
