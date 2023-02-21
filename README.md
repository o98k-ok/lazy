
## lazy golang package 

### 1. alfred

Just a alfred toolkit, can be used to create alfred plugin quikcly. <br />

e.g.
```golang
func entry() {
	cli := alfred.NewApp("new doc search plugin")
	cli.Bind("query", func(s []string) {
		env, err := alfred.GetFlowEnv()
		if err != nil {
			alfred.ErrItems("alfred get envs failed", err).Show()
			return
		}

		cli := doc.NewLark().CustomSession(env.GetAsString("session", "")).WithPage(0, env.GetAsInt("count", 9))
		params := strings.Join(s, " ")
		entities, err := cli.Query(params)
		if err != nil {
			alfred.ErrItems("query lark failed", err)
			return
		}

		msg := alfred.NewItems()
		for _, entity := range entities {
			item := alfred.NewItem(title(entity), intro(entity), entity.Url)
			item.Extra = entity.ViewTS
			msg.Append(item)
		}

		msg.Order(func(l, r *alfred.Item) bool {
			return l.Extra.(uint32) > r.Extra.(uint32)
		})
		msg.Show()
	})

	if err := cli.Run(os.Args); err != nil {
		alfred.ErrItems("alfred run failed", err).Show()
		return
	}
}
```

More you can get from https://github.com/o98k-ok/awesome-alfred-workflow

### 2. cache

A simple cache interface, just imp file cache.

### 3. format

Often you need format json, so this package would help you.
1. format with spec char;
2. format with custom color;


### 4. mac

Some mac interface, you can read/write plist easily.


### 5. route

A simple pack for iris, you can do following via it:
1. Custom your core handler
2. Custom any compenents of route
3. Provide a simple version for generate api doc

```golang
type Level1 struct {
		Age *string `schema:"age" fake:"{word}"`
	}
	type Level2 struct {
		Name  string  `json:"name"  fake:"{firstname}"`
		Level *Level1 `json:"level"`
	}

	var l1 Level1
	var l2 Level2
	gofakeit.Struct(&l2)
	gofakeit.Struct(&l1)
	elems := Elems{
		Method: "GET",
		URI:    "/api/v1/users",
		Req:    l1,
		Resp:   l2,
	}

	fmt.Println(GenerateAPIDoc(elems))
```

will get result:

```markdown
## 1. 接口简介

|  类型  |     信息      | 备注 |
|--------|---------------|------|
| URI    | /api/v1/users |      |
| METHOD | GET           |      |

## 2. 参数信息

请求数据类型为: Level1
| 字段名称 | 字段类型 | 字段含义 | 是否必要 | 备注 |
|----------|----------|----------|----------|------|
| age      | string   |          | NO       |      |

## 3. 返回信息

返回数据类型为: Level2
| 字段名称 | 字段类型 | 字段含义 | 备注 |
|----------|----------|----------|------|
| name     | string   |          |      |
| level    | Level1   |          |      |

## 4. 请求示例


```json
GET /api/v1/users?age=what

{
    "name": "Maryse",
    "level": {
        "Age": "should"
    }
}
```

```
