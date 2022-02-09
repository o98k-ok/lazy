package main

import (
	"encoding/json"
	"fmt"
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

	nodes, _ := host.NewNodeGraph(structed.Chains, structed.Nodes)
	fmt.Println(nodes.ListNodeByTags([]string{"dev"})[0].IP)
	fmt.Println(nodes.ListNodeByKey("10")[0].IP)
	for _, node := range nodes.GenNodeRelation(nodes.ListNodeByKey("10")[0]) {
		fmt.Println(node.IP)
	}
}
