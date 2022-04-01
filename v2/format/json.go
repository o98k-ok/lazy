package format

import (
	"io"
	"reflect"
	"strconv"
	"strings"
)

const (
	Comma     = ","
	ArrStart  = "["
	ArrEnd    = "]"
	ObjStart  = "{"
	ObjEnd    = "}"
	Ln        = "\n"
	Null      = "null"
	Space     = " "
	StringSeg = "\""
	colon     = ": "
)

type JsonFormatter struct {
	w      io.Writer
	Indent string

	numberColor func(a ...interface{})
	mapKeyColor func(a ...interface{})
	boolColor   func(a ...interface{})
	nullColor   func(a ...interface{})
	stringColor func(a ...interface{})
	noColor     func(a ...interface{})
}

func (j *JsonFormatter) processString(val reflect.Value) {
	j.noColor(StringSeg)
	j.stringColor(val.String())
	j.noColor(StringSeg)
}

func (j *JsonFormatter) processMapKey(val reflect.Value) {
	j.noColor(StringSeg)
	j.mapKeyColor(val.String())
	j.noColor(StringSeg)
}

func (j *JsonFormatter) processInt(val reflect.Value) {
	j.numberColor(strconv.FormatInt(val.Int(), 10))
}

func (j *JsonFormatter) processFloat(val reflect.Value, size int) {
	j.numberColor(strconv.FormatFloat(val.Float(), 'f', -1, size))
}

func (j *JsonFormatter) processFloat64(val reflect.Value) {
	j.processFloat(val, 64)
}

func (j *JsonFormatter) processFloat32(val reflect.Value) {
	j.processFloat(val, 32)
}

func (j *JsonFormatter) processSlice(val reflect.Value, depth int) {
	j.noColor(ArrStart)
	j.noColor(Ln)

	for i := 0; i < val.Len(); i++ {
		j.noColor(strings.Repeat(j.Indent, depth+1))
		j.process(val.Index(i), depth+1)
		if i != val.Len()-1 {
			j.noColor(Comma)
		}
		j.noColor(Ln)
	}

	j.noColor(strings.Repeat(j.Indent, depth))
	j.noColor(ArrEnd)
}

func (j *JsonFormatter) processMap(val reflect.Value, depth int) {
	j.noColor(ObjStart)
	j.noColor(Ln)

	keys := val.MapKeys()
	for i := 0; i < len(keys); i++ {
		j.noColor(strings.Repeat(j.Indent, depth+1))
		j.processMapKey(keys[i])
		j.noColor(colon)
		j.process(val.MapIndex(keys[i]), depth+1)
		if i != len(keys)-1 {
			j.noColor(Comma)
		}
		j.noColor(Ln)
	}
	j.noColor(strings.Repeat(j.Indent, depth))
	j.noColor(ObjEnd)
}

func (j *JsonFormatter) process(val reflect.Value, depth int) {
	switch val.Kind() {
	case reflect.Map:
		j.processMap(val, depth)
	case reflect.Slice, reflect.Array:
		j.processSlice(val, depth)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		j.processInt(val)
	case reflect.String:
		j.processString(val)
	case reflect.Float64:
		j.processFloat64(val)
	case reflect.Float32:
		j.processFloat32(val)
	case reflect.Bool:
		j.boolColor(strconv.FormatBool(val.Bool()))
	case reflect.Invalid:
		j.nullColor(Null)
	}
}

func (j *JsonFormatter) Encode(obj any) error {
	val := reflect.ValueOf(obj)
	j.process(val, 0)
	return nil
}

func NewEncoder(writer io.Writer) *JsonFormatter {
	formatter := &JsonFormatter{
		w:      writer,
		Indent: strings.Repeat(Space, 4),
	}

	formatter.initColor()
	return formatter
}

func (j *JsonFormatter) WithIndent(indent string) *JsonFormatter {
	j.Indent = indent
	return j
}
