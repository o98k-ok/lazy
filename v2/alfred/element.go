package alfred

type Icon struct {
	Path string `json:"path,omitempty"`
}

type Item struct {
	Arg      string `json:"arg,omitempty"`
	Title    string `json:"title"`
	SubTitle string `json:"subtitle,omitempty"`
	Icon     *Icon  `json:"icon,omitempty"`
	Extra    interface{}
}

func NewItem(title, subtitle, arg string) *Item {
	return &Item{
		Arg:      arg,
		Title:    title,
		SubTitle: subtitle,
	}
}

func (i *Item) WithIcon(path string) {
	i.Icon.Path = path
}
