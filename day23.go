package main

import (
	"fmt"
	"slices"
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
	slices.Sort(nodesList)
	maxClique := findMaximumClique(make([]string, 0), nodesList, edges)
	fmt.Println(strings.Join(maxClique, ","))
}

func findMaximumClique(current []string, nodes []string, edges map[string]bool) []string {
	maxClique := current
nodeCycle:
	for i, node := range nodes {
		for _, cur := range current {
			if !edges[cur+"-"+node] {
				continue nodeCycle
			}
		}
		next := slices.Concat(current, []string{node})
		nextClique := findMaximumClique(next, nodes[i+1:], edges)
		if len(nextClique) > len(maxClique) {
			maxClique = nextClique
		}
	}
	return maxClique
}
