package route

import (
	"fmt"
	"reflect"

	"github.com/fatih/structtag"
)

type TypeDetail struct {
	TypeName   string
	Anonymous  bool
	DeepInType reflect.Type
}

type handle func(tpe reflect.Type)

func Jump(real reflect.Type, pointerH handle, sliceH handle) reflect.Type {
	for {
		if IsPointer(real) {
			if pointerH != nil {
				pointerH(real)
			}
			real = real.Elem()
			continue
		}

		if IsSlice(real) {
			if sliceH != nil {
				sliceH(real)
			}
			real = real.Elem()
			continue
		}
		return real
	}
}

func IsPointer(real reflect.Type) bool {
	return real.Kind() == reflect.Pointer
}

func IsSlice(real reflect.Type) bool {
	return real.Kind() == reflect.Array || real.Kind() == reflect.Slice
}

func GetFieldInfo(field reflect.StructField) TypeDetail {
	var prefix string
	real := Jump(field.Type, nil, func(tpe reflect.Type) { prefix = "[]" + prefix })

	res := TypeDetail{
		TypeName:   prefix + real.Name(),
		Anonymous:  field.Anonymous,
		DeepInType: real,
	}
	return res
}

type TypeConfig struct {
	NameTag     string
	DescTag     string
	ValidateTag string
}

func (t TypeConfig) needValid() bool {
	return len(t.ValidateTag) > 0
}

func (t TypeConfig) getValidateTag(tags *structtag.Tags) string {
	if t.ValidateTag == "" {
		return ""
	}
	tag, err := tags.Get(t.ValidateTag)
	if err == nil && tag.Name == "required" {
		return "YES"
	} else {
		return "NO"
	}
}

func (t TypeConfig) getDescTag(tags *structtag.Tags) string {
	tag, err := tags.Get(t.DescTag)
	if err != nil || tag.Name == "-" {
		return ""
	} else {
		return tag.Name
	}
}

func (t TypeConfig) getNameTag(tags *structtag.Tags, indent string) string {
	tag, err := tags.Get(t.NameTag)
	if err != nil || tag.Name == "-" {
		return ""
	} else {
		return indent + tag.Name
	}
}

func (t TypeConfig) ParseStruct(tpe reflect.Type, namePrefix string) [][]string {
	var res [][]string

	for i := 0; i < tpe.NumField(); i++ {
		field := tpe.Field(i)
		info := GetFieldInfo(field)

		// only support Anonymous && struct
		if info.Anonymous && info.DeepInType.Kind() == reflect.Struct {
			res = append(res, t.ParseStruct(info.DeepInType, namePrefix)...)
			continue
		}

		var fields []string
		tags, err := structtag.Parse(string(field.Tag))
		if err != nil {
			continue
		}

		tagname := t.getNameTag(tags, namePrefix)
		if len(tagname) == 0 {
			continue
		}
		fields = append(fields, tagname)
		fields = append(fields, info.TypeName)
		fields = append(fields, t.getDescTag(tags))
		if t.needValid() {
			fields = append(fields, t.getValidateTag(tags))
		}
		fields = append(fields, "")
		res = append(res, fields)

		if info.DeepInType.Kind() == reflect.Struct {
			res = append(res, t.ParseStruct(info.DeepInType, tagname+"."+namePrefix)...)
		}
	}
	return res
}

func (t TypeConfig) GenerateTable(tpe reflect.Type) ([][]string, string) {
	var prefix string
	real := Jump(tpe, nil, func(tpe reflect.Type) { prefix = "[]" + prefix })

	switch real.Kind() {
	case reflect.Struct:
		return t.ParseStruct(real, ""), prefix + real.Name()
	case reflect.Complex128, reflect.Complex64, reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.UnsafePointer:
		fmt.Println("unsupport type parse")
		return nil, ""
	default:
		return nil, prefix + real.Name()
	}
}
