package alfred

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Items struct {
	Items []*Item `json:"items"`
}

func (i *Items) Len() int {
	return len(i.Items)
}

func NewItems() *Items {
	return &Items{Items: make([]*Item, 0)}
}

func (i *Items) Append(item *Item) *Items {
	i.Items = append(i.Items, item)
	return i
}

func (i *Items) Encode() string {
	dat, _ := json.Marshal(*i)
	return string(dat)
}

func ErrItems(title string, err error) *Items {
	return errItemsWithLog(title, err, os.Stderr)
}

func errItemsWithLog(title string, err error, writer io.Writer) *Items {
	log(writer, "%s err %v\n", title, err)
	return NewItems().Append(NewItem(title, err.Error(), ""))
}

func EmptyItems() *Items {
	return ErrItems("404", errors.New("empty result"))
}
