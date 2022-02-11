package host

import "testing"


var nodes = []Node{
	{
		Name: "server1",
		Tags: []string{"test", "001"},
		IP:   "10.10.10",
		Port: "360000",
		User: "root",
	},
	{
		Name: "server2",
		Tags: []string{"test", "002"},
		IP:   "10.10.10.10",
		Port: "22",
		User: "root",
	},
	{
		Name:   "server3",
		Tags:   []string{"test", "003"},
		IP:     "10.10.10.20",
		Port:   "22",
		User:   "root",
		Depend: "server1",
	},
}

func TestNodeGraph_GenNodeRelation(t *testing.T) {
	graph, err := NewNodeGraph(nodes)
	if err != nil {
		t.Fatal(err)
	}

	ns := graph.ListNodeByKey("server3")
	if len(ns) != 1 {
		t.Fatal()
	}

	res := graph.GenNodeRelation(ns[0])
	if len(res) != 2 {
		t.Fatal()
	}
}
