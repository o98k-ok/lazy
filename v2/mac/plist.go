package mac

import (
	"fmt"
	"howett.net/plist"
	"io"
)

func DefaultsRead(reader io.ReadSeeker, attrs []string) (interface{}, error) {
	var data map[string]interface{}

	decoder := plist.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return getAttr(attrs, data)
}

func getAttr(keys []string, attr interface{}) (interface{}, error) {
	for _, c := range keys {
		if attr == nil {
			return nil, fmt.Errorf("no such keys %v", keys)
		}

		val, ok := attr.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("cannot parse keys %v", keys)
		}

		attr = val[c]
	}
	return attr, nil
}
