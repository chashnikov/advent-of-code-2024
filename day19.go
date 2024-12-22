package main

import (
	"fmt"
	"strings"
)

func day19(fileName string) {
	lines := ReadLines(fileName)
	patterns := strings.Split(lines[0], ", ")
	fmt.Println(Count(lines[2:], func(design string) bool {
		return isTowelDesignPossible(design, patterns)
	}))
}

func isTowelDesignPossible(design string, patterns []string) bool {
	reachable := make([]bool, len(design)+1)
	reachable[0] = true
	for i := 0; i < len(design); i++ {
		if !reachable[i] {
			continue
		}
		for _, pattern := range patterns {
			if i+len(pattern) <= len(design) && design[i:i+len(pattern)] == pattern {
				reachable[i+len(pattern)] = true
			}
		}
	}
	return reachable[len(design)]
}
