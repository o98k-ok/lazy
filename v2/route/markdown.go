package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/muesli/marky"
	"github.com/o98k-ok/schema"
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
	Fn     func(interface{}) interface{}
}

type FrontConf struct {
	Title     string `json:"title"`
	DataIndex string `json:"dataIndex"`
	ValueType string `json:"valueType,omitempty"`
	Required  *bool  `json:"required,omitempty"`
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

	var req TypeConfig
	if elem.Method == http.MethodGet {
		req = TypeConfig{
			NameTag:     "schema",
			DescTag:     "desc",
			ValidateTag: "validate",
		}
	} else {
		req = TypeConfig{
			NameTag:     "json",
			DescTag:     "desc",
			ValidateTag: "validate",
		}
	}
	requestTable := req.GenerateTable(reflect.TypeOf(elem.Req))
	for t, table := range requestTable {
		doc.Add(marky.Text{Text: fmt.Sprintf("请求数据类型为: %s\n", t)})
		doc.Add(NewMarkyTable(reqHeader, table))
		doc.Add(&marky.BlockElement{})
	}

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "3. 返回信息",
	})

	resp := TypeConfig{
		NameTag: "json",
		DescTag: "desc",
	}
	respHeader := []string{"字段名称", "字段类型", "字段含义", "备注"}
	responseTable := resp.GenerateTable(reflect.TypeOf(elem.Resp))
	for t, table := range responseTable {
		doc.Add(marky.Text{Text: fmt.Sprintf("返回数据类型为: %s\n", t)})
		doc.Add(NewMarkyTable(respHeader, table))
		doc.Add(&marky.BlockElement{})
	}

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "4. 参数配置",
	})

	if elem.Method == http.MethodPost || elem.Method == http.MethodPut {
		// using request
		for key, table := range requestTable {
			var field []FrontConf
			for _, t := range table {
				reqiured := strings.EqualFold(t[3], "YES")
				field = append(field, FrontConf{
					Title:     t[2],
					DataIndex: t[0],
					Required:  &reqiured,
				})
			}
			doc.Add(marky.Text{Text: fmt.Sprintf("返回类型: %s\n", key)})
			doc.Add(marky.Code{
				Source:   FormatJson(field) + "\n",
				Language: "json",
			})
		}
	} else {
		for key, table := range responseTable {
			var field []FrontConf
			for _, t := range table {
				field = append(field, FrontConf{
					Title:     t[2],
					DataIndex: t[0],
				})
			}
			doc.Add(marky.Text{Text: fmt.Sprintf("返回类型: %s\n", key)})
			doc.Add(marky.Code{
				Source:   FormatJson(field) + "\n",
				Language: "json",
			})
		}
	}
	doc.Add(&marky.BlockElement{})

	doc.Add(marky.Heading{
		Level:   2,
		Caption: "5. 请求示例",
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

	if elem.Method == http.MethodGet {
		var mm map[string][]string = make(map[string][]string)
		schema.NewEncoder().Encode(elem.Req, mm)

		u := url.Values{}
		for k, v := range mm {
			u.Add(k, strings.Join(v, ","))
		}

		if len(u) > 0 {
			code.WriteString("?")
			code.WriteString(u.Encode())
		}
	} else {
		code.WriteString("\n\n")
		code.WriteString(FormatJson(elem.Req))
	}

	code.WriteString("\n\n")
	code.WriteString(FormatJson(elem.Fn(elem.Resp)))
	code.WriteString("\n")
	return code.String()
}

func FormatJson(en interface{}) string {
	dat, _ := json.MarshalIndent(en, "", "  ")
	return string(dat)
}
