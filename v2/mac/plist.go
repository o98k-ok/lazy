package mac

import (
	"github.com/o98k-ok/lazy/v2/collection"
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

	return collection.GetAttr(attrs, data)
}
