package route

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/fatih/structtag"
	"github.com/gorilla/schema"
	"github.com/muesli/marky"
	"github.com/o98k-ok/lazy/v2/format"
	"github.com/olekukonko/tablewriter"
)

type MarkyTable struct {
	writer bytes.Buffer
}

func NewMarkyTable(header []string, data [][]string) MarkyTable {
	t := MarkyTable{}
	table := tablewriter.NewWriter(&t.writer)
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return t
}

func (m MarkyTable) String() string {
	return m.writer.String()
}

type Elems struct {
	Method string
	URI    string
	Req    interface{}
	Resp   interface{}
}

func GenerateAPIDoc(elem Elems) (string, error) {
	doc := marky.NewDocument()

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "1. 接口简介",
	})
	header := []string{"类型", "信息", "备注"}
	doc.Add(NewMarkyTable(header, [][]string{
		{"URI", elem.URI, ""},
		{"METHOD", elem.Method, ""},
	}))
	doc.Add(&marky.BlockElement{})

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "2. 参数信息",
	})
	reqHeader := []string{"字段名称", "字段类型", "字段含义", "是否必要", "备注"}
	if elem.Method == http.MethodGet {
		doc.Add(NewMarkyTable(reqHeader, RequestTable(elem.Req, "schema")))
	} else {
		doc.Add(NewMarkyTable(reqHeader, RequestTable(elem.Req, "json")))
	}
	doc.Add(&marky.BlockElement{})

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "3. 返回信息",
	})
	respHeader := []string{"字段名称", "字段类型", "字段含义", "备注"}
	doc.Add(NewMarkyTable(respHeader, ResponseTable(elem.Resp)))
	doc.Add(&marky.BlockElement{})

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "4. 请求示例",
	})

	doc.Add(marky.Code{
		Source:   FormatDemoCode(elem),
		Language: "json",
	})
	doc.Add(&marky.BlockElement{})
	return doc.String(), nil
}

func FormatDemoCode(elem Elems) string {
	code := strings.Builder{}
	code.WriteString(elem.Method)
	code.WriteString(" ")
	code.WriteString(elem.URI)

	var mm map[string][]string = make(map[string][]string)
	schema.NewEncoder().Encode(elem.Req, mm)

	u := url.Values{}
	for k, v := range mm {
		u.Add(k, strings.Join(v, ","))
	}
	code.WriteString("?")
	code.WriteString(u.Encode())

	code.WriteString("\n\n")
	code.WriteString(FormatJson(elem.Req))
	code.WriteString("\n")
	code.WriteString("\n")
	code.WriteString(FormatJson(elem.Resp))
	code.WriteString("\n")
	return code.String()
}

func FormatJson(en interface{}) string {
	dat, _ := json.Marshal(en)
	var tmp map[string]interface{}
	json.Unmarshal(dat, &tmp)

	res := &bytes.Buffer{}
	format.NewEncoder(res).Encode(tmp)
	return res.String()
}

// RequestTable []string{"字段名称", "字段类型", "字段含义", "是否必要", "备注"},
func RequestTable(req interface{}, nameTag string) [][]string {
	descTag, validTag := "desc", "validate"
	tpe := reflect.TypeOf(req)

	var res [][]string
	for i := 0; i < tpe.NumField(); i++ {
		var fields []string
		tags, err := structtag.Parse(string(tpe.Field(i).Tag))
		if err != nil {
			continue
		}

		tag, err := tags.Get(nameTag)
		if err != nil || tag.Name == "-" {
			continue
		}
		fields = append(fields, tag.Name)
		fields = append(fields, tpe.Field(i).Type.Name())

		tag, err = tags.Get(descTag)
		if err != nil || tag.Name == "-" {
			fields = append(fields, "")
		} else {
			fields = append(fields, tag.Name)
		}

		tag, err = tags.Get(validTag)
		if err == nil && tag.Name == "required" {
			fields = append(fields, "YES")
		} else {
			fields = append(fields, "NO")
		}

		fields = append(fields, "")
		res = append(res, fields)
	}
	return res
}

func ResponseTable(req interface{}) [][]string {
	nameTag, descTag := "json", "desc"
	tpe := reflect.TypeOf(req)

	var res [][]string
	for i := 0; i < tpe.NumField(); i++ {
		var fields []string
		tags, err := structtag.Parse(string(tpe.Field(i).Tag))
		if err != nil {
			continue
		}

		tag, err := tags.Get(nameTag)
		if err != nil || tag.Name == "-" {
			continue
		}
		fields = append(fields, tag.Name)
		fields = append(fields, tpe.Field(i).Type.Name())

		tag, err = tags.Get(descTag)
		if err != nil || tag.Name == "-" {
			fields = append(fields, "")
		} else {
			fields = append(fields, tag.Name)
		}

		fields = append(fields, "")
		res = append(res, fields)
	}
	return res
}
