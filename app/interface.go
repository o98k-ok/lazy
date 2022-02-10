package app

import (
	"fmt"
	"github.com/o98k-ok/lazy/host"
	"github.com/o98k-ok/lazy/utils"
	"strings"
)

var Nodes *host.NodeGraph
const IconPath = "../conf/lazy.jpeg"

func InitApp(chains [][]string, nodes []host.Node)  error {
	var err error
	Nodes, err = host.NewNodeGraph(chains, nodes)
	return err
}

// FilterByKeys
// filter by key1 or key2 or ....
func FilterByKeys(keys []string) *Items {
	resItems := NewItems()

	allnodes := make([]*host.Node, 0)
	for _, key := range keys {
		allnodes = append(allnodes, Nodes.ListNodeByKey(key)...)
	}

	for _, node := range allnodes {
		resItems.Append(&Item{
			Arg:      LoginLine(node),
			Title:    node.Name,
			SubTitle: fmt.Sprintf("%s-[%s]", node.IP, strings.Join(node.Tags, "|")),
			Icon:     Icon{Path: IconPath},
		})
	}
	return resItems
}

// FilterByTags
// filter by tag1 and tag2 and ...
func FilterByTags(tags []string) *Items {
	resItems := NewItems()
	for _, node := range Nodes.ListNodeByTags(tags) {
		resItems.Append(&Item{
			Arg:      LoginLine(node),
			Title:    node.Name,
			SubTitle: fmt.Sprintf("%s-[%s]", node.IP, strings.Join(node.Tags, "|")),
			Icon:     Icon{Path: IconPath},
		})
	}
	return resItems
}

// LoginLine
// title#cmd1;cmd2;...
func LoginLine(node *host.Node) string {
	dep := Nodes.GenNodeRelation(node)

	res := make([]string, 0)
	for _, n := range dep {
		// login command
		var cmd string
		switch  {
		case !utils.Empty(n.User) && !utils.Empty(n.Port):
			cmd = fmt.Sprintf("ssh %s@%s -p%s", n.User, n.IP, n.Port)
		case !utils.Empty(n.User) && utils.Empty(n.Port):
			cmd = fmt.Sprintf("ssh %s@%s", n.User, n.IP)
		case utils.Empty(n.User) && !utils.Empty(n.Port):
			cmd = fmt.Sprintf("ssh %s -p%s", n.IP, n.Port)
		case utils.Empty(n.User) && utils.Empty(n.Port):
			cmd = fmt.Sprintf("ssh %s", n.IP)
		}

		if !utils.Empty(n.Passwd) {
			cmd += fmt.Sprintf(";%s", n.Passwd)
		}

		res = append(res, cmd)
	}
	return fmt.Sprintf("%s#%s", node.Name, strings.Join(res, ";"))
}