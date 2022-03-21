package mac

import (
	"fmt"
	"howett.net/plist"
	"io"
)

func DefaultsWrite(reader io.ReadSeeker, attrs []string, value interface{}, writer io.Writer) error {
	var data map[string]interface{}

	decoder := plist.NewDecoder(reader)
	err := decoder.Decode(&data)
	if err != nil {
		return err
	}

	for i := len(attrs) - 1; i > 0; i-- {
		node := make(map[string]interface{})
		node[attrs[i]] = value
		value = node
	}
	data[attrs[0]] = value

	format := decoder.Format
	return plist.NewEncoderForFormat(writer, format).Encode(data)
}

func DefaultsRead(reader io.ReadSeeker, attrs []string) (interface{}, error) {
	var data map[string]interface{}

	err := plist.NewDecoder(reader).Decode(&data)
	if err != nil {
		return nil, err
	}

	return getAttr(attrs, data)
}

func getAttr(keys []string, attr interface{}) (interface{}, error) {
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
