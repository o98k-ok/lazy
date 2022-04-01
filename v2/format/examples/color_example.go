package main

import (
	"fmt"
	"github.com/o98k-ok/lazy/v2/format"
	"os"
)

func main() {
	mm := map[string]int{
		"shadow":  100,
		"o98k-ok": 100,
	}

	encoder := format.NewEncoder(os.Stdout)
	err := encoder.Encode(mm)
	if err != nil {
		fmt.Println(err)
	}

	deepMap := map[string]map[string][]int{
		"shadow": {
			"math":    []int{10, 20, 30, 40},
			"chinese": []int{100, 120, 130, 140},
		},
		"hulk": {
			"math":    []int{},
			"chinese": nil,
		},
	}
	err = encoder.Encode(deepMap)
	if err != nil {
		fmt.Println(err)
	}

	encoder.NumberColor(format.Green).MapKeyColor(format.Cyan)
	err = encoder.Encode(mm)
	if err != nil {
		fmt.Println(err)
	}

	encoder.DisableColor()
	err = encoder.Encode(mm)
	if err != nil {
		fmt.Println(err)
	}
}
