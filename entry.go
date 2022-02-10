package main

import (
	"encoding/json"
	"fmt"
	"github.com/o98k-ok/lazy/app"
	"github.com/o98k-ok/lazy/host"
	"io/ioutil"
)

func main() {
	d, err := ioutil.ReadFile("./conf/conf.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	structed := struct {
		Chains [][]string  `json:"chains"`
		Nodes  []host.Node `json:"nodes"`
	}{}

	err = json.Unmarshal(d, &structed)
	if err != nil {
		fmt.Println(err)
		return
	}

	app.InitApp(structed.Chains, structed.Nodes)
	fmt.Println(app.FilterByKeys([]string{"com"}).Encode())
	fmt.Println(app.FilterByTags([]string{"jum"}).Encode())
}
