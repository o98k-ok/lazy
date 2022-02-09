package host

import (
	"github.com/o98k-ok/lazy/utils"
	"strings"
)

type Node struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
	Type string   `json:"type"`
	// Nid string

	IP     string `json:"ip"`
	Port   string `json:"port"`
	User   string `json:"user"`
	Passwd string `json:"passwd"`

	Extra string `json:"extra"`
}

type NodeGraph struct {
	// all nodes type relations
	Types TypeDependency
	// all nodes list: key=ip
	Nodes map[string]*Node
	// key: node one tag
	Tags map[string][]*Node
	// key: type name
	TypeNodes map[string][]*Node
}

func NewNodeGraph(chains [][]string, nodes []Node) (*NodeGraph, error) {
	resGraph := NodeGraph{
		Types:     NewTypeDependency(chains),
		Nodes:     make(map[string]*Node),
		Tags:      make(map[string][]*Node),
		TypeNodes: make(map[string][]*Node),
	}

	for i, elem := range nodes {
		resGraph.Nodes[elem.IP] = &nodes[i]
		if _, ok := resGraph.TypeNodes[elem.Type]; !ok {
			resGraph.TypeNodes[elem.Type] = make([]*Node, 0)
		}
		resGraph.TypeNodes[elem.Type] = append(resGraph.TypeNodes[elem.Type], &nodes[i])

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

	mtype := node.Type
	relationLine, ok := ng.Types[mtype]
	if !ok {
		return res
	}

	for _, rela := range relationLine {
		res = append(res, GetFirstNode(ng.TypeNodes[rela.Name]))
	}

	return res
}

// TODO here you can expand your node info
func GetFirstNode(nodes []*Node) *Node {
	if len(nodes) <= 0 {
		return nil
	}

	return nodes[0]
}
