package host

import (
	"github.com/o98k-ok/lazy/utils"
	"strings"
)

type Node struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`

	IP     string `json:"ip"`
	Port   string `json:"port"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`
	Depend string `json:"depend"`

	Extra string `json:"extra"`
}

type NodeGraph struct {
	// all nodes list: key=name
	Nodes map[string]*Node
	// key: node one tag
	Tags map[string][]*Node
}

func NewNodeGraph(nodes []Node) (*NodeGraph, error) {
	resGraph := NodeGraph{
		Nodes:     make(map[string]*Node),
		Tags:      make(map[string][]*Node),
	}

	for i, elem := range nodes {
		resGraph.Nodes[elem.Name] = &nodes[i]

		for _, t := range elem.Tags {
			if _, ok := resGraph.Tags[t]; !ok {
				resGraph.Tags[t] = make([]*Node, 0)
			}
			resGraph.Tags[t] = append(resGraph.Tags[t], &nodes[i])
		}
	}
	return &resGraph, nil
}

func LikeIn(match, tag string) bool {
	return strings.Contains(tag, match)
}

func MatchByTag(tag string, tagInfo map[string][]*Node) []*Node {
	res := make([]*Node, 0)
	for key, value := range tagInfo {
		if !strings.Contains(key, tag) {
			continue
		}
		res = append(res, value...)
	}
	return res
}

func (ng *NodeGraph) ListNodeByTags(tags []string) []*Node {
	if len(tags) <= 0 {
		return make([]*Node, 0)
	}

	resSet := utils.ToSet(MatchByTag(tags[0], ng.Tags))
	for i := 1; i < len(tags); i++ {
		set := utils.ToSet(MatchByTag(tags[i], ng.Tags))
		resSet = utils.AND(resSet, set)
	}

	res := make([]*Node, 0)
	for node, _ := range resSet {
		res = append(res, node.(*Node))
	}
	return res
}

// ListNodeByKey
//  @param key
// 		substr of Node Name or Node IP
//  @return []*Node
// 		node list
func (ng *NodeGraph) ListNodeByKey(key string) []*Node {
	res := make([]*Node, 0)
	for name, node := range ng.Nodes {
		if !strings.Contains(name, key) && !strings.Contains(node.IP, key) {
			continue
		}
		res = append(res, node)
	}
	return res
}

func (ng *NodeGraph) GenNodeRelation(node *Node) []*Node {
	res := make([]*Node, 0)

	for  {
		// reverse order
		res = append([]*Node{node}, res...)
		dependKey := node.Depend
		if utils.Empty(dependKey) {
			break
		}

		var ok bool
		if node, ok = ng.Nodes[dependKey]; !ok {
			return make([]*Node, 0)
		}
	}

	return res
}
