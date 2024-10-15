package alfred

type Icon struct {
	Path string `json:"path,omitempty"`
}

type Item struct {
	Arg       string      `json:"arg,omitempty"`
	Title     string      `json:"title"`
	SubTitle  string      `json:"subtitle,omitempty"`
	Icon      *Icon       `json:"icon,omitempty"`
	Uid       string      `json:"uid,omitempty"`
	Variables Variables   `json:"variables,omitempty"`
	Extra     interface{} `json:"-"`
}

type Variables map[string]string

func NewItem(title, subtitle, arg string) *Item {
	return &Item{
		Arg:      arg,
		Title:    title,
		SubTitle: subtitle,
	}
}

func (i *Item) WithVariable(key, val string) *Item {
	if i.Variables == nil {
		i.Variables = make(Variables)
	}
	i.Variables[key] = val
	return i
}

func (i *Item) WithUid(uid string) *Item {
	i.Uid = uid
	return i
}

func (i *Item) WithExtra(extra interface{}) *Item {
	i.Extra = extra
	return i
}

func (i *Item) WithIcon(path string) *Item {
	i.Icon = &Icon{Path: path}
	return i
}
