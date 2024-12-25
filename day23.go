package main

import (
	"fmt"
	"strings"
)

func day23(fileName string) {
	lines := ReadLines(fileName)
	edges := make(map[string]bool)
	nodes := make(map[string]bool)
	nodesList := make([]string, 0)
	for _, line := range lines {
		ends := strings.Split(line, "-")
		edges[ends[0]+"-"+ends[1]] = true
		edges[ends[1]+"-"+ends[0]] = true
		if !nodes[ends[0]] {
			nodes[ends[0]] = true
			nodesList = append(nodesList, ends[0])
		}
		if !nodes[ends[1]] {
			nodes[ends[1]] = true
			nodesList = append(nodesList, ends[1])
		}
	}
	res := 0
	for i := 0; i < len(nodesList); i++ {
		n1 := nodesList[i]
		for j := i + 1; j < len(nodesList); j++ {
			n2 := nodesList[j]
			if !edges[n1+"-"+n2] {
				continue
			}
			for k := j + 1; k < len(nodesList); k++ {
				n3 := nodesList[k]
				if !strings.HasPrefix(n1, "t") && !strings.HasPrefix(n2, "t") && !strings.HasPrefix(n3, "t") {
					continue
				}
				if !edges[n1+"-"+n3] || !edges[n2+"-"+n3] {
					continue
				}
				res++
			}
		}
	}
	fmt.Println(res)
}
