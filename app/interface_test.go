package app

import (
	"github.com/o98k-ok/lazy/host"
	"testing"
)

var nodes = []host.Node{
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

func TestFilterByTags(t *testing.T) {
	if err := InitApp(nodes); err != nil {
		t.Fatal(err)
	}

	res := FilterByTags([]string{"test"})
	if len(res.Items) != 3 {
		t.Fail()
	}
}

func TestFilterByKeys(t *testing.T) {
	if err := InitApp(nodes); err != nil {
		t.Fatal(err)
	}

	res := FilterByKeys([]string{"20"})
	if len(res.Items) != 2 {
		t.Fail()
	}
}