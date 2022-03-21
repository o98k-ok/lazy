package mac

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

var dat = []byte{98, 112, 108, 105, 115, 116, 48, 48, 214, 1, 2, 3, 4, 5, 6, 7, 7, 9, 10, 11, 12, 95, 16, 15, 115, 97, 118, 101, 45, 115, 101, 108, 101, 99, 116, 105, 111, 110, 115, 94, 115, 104, 111, 119, 45, 116, 104, 117, 109, 98, 110, 97, 105, 108, 85, 115, 116, 121, 108, 101, 95, 16, 20, 108, 97, 115, 116, 45, 97, 110, 97, 108, 121, 116, 105, 99, 115, 45, 115, 116, 97, 109, 112, 85, 118, 105, 100, 101, 111, 88, 108, 111, 99, 97, 116, 105, 111, 110, 8, 8, 89, 115, 101, 108, 101, 99, 116, 105, 111, 110, 35, 65, 195, 240, 220, 154, 205, 41, 128, 9, 95, 16, 34, 47, 85, 115, 101, 114, 115, 47, 115, 104, 97, 100, 111, 119, 47, 68, 111, 99, 117, 109, 101, 110, 116, 115, 47, 115, 99, 114, 101, 101, 110, 115, 104, 111, 116, 8, 21, 39, 54, 60, 83, 89, 98, 99, 100, 110, 119, 120, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0, 0, 0, 0, 13, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 157}

func TestDefaultsRead(t *testing.T) {
	reader := bytes.NewReader(dat)

	t.Run("test error cases", func(t *testing.T) {
		_, _ = reader.Seek(0, io.SeekStart)
		_, err := DefaultsRead(reader, []string{"locations"})
		assert.Error(t, err)
	})

	t.Run("test empty content cases", func(t *testing.T) {
		reader.Seek(0, io.SeekEnd)
		_, err := DefaultsRead(reader, []string{"location"})
		assert.Error(t, err)
	})

	t.Run("test normal cases", func(t *testing.T) {
		_, _ = reader.Seek(0, io.SeekStart)
		attr, err := DefaultsRead(reader, []string{"location"})
		assert.NoError(t, err)

		expect := "/Users/shadow/Documents/screenshot"
		assert.Equal(t, expect, attr)
	})

	t.Run("test inner cases", func(t *testing.T) {
		res, err := getAttr([]string{"location"}, nil)
		assert.NoError(t, err)
		assert.Equal(t, nil, res)
	})
}

func TestDefaultsWrite(t *testing.T) {
	reader := bytes.NewReader(dat)
	writer := &bytes.Buffer{}

	attrs := []string{"location"}
	expect := "/Users/shadow/Desktop"
	err := DefaultsWrite(reader, attrs, expect, writer)
	assert.NoError(t, err)

	t.Run("test changed fields", func(t *testing.T) {
		newReader := bytes.NewReader(writer.Bytes())
		res, err := DefaultsRead(newReader, attrs)
		assert.NoError(t, err)
		assert.Equal(t, expect, res)
	})

	t.Run("test unchanged fields", func(t *testing.T) {
		attr := []string{"style"}
		reader.Seek(0, io.SeekStart)
		before, err := DefaultsRead(reader, attr)
		assert.NoError(t, err)

		newReader := bytes.NewReader(writer.Bytes())
		after, err := DefaultsRead(newReader, attr)
		assert.NoError(t, err)

		assert.Equal(t, before, after)
	})
}
