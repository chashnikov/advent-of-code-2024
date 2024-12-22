package main

import (
	"fmt"
	"strings"
)

func day19(fileName string) {
	lines := ReadLines(fileName)
	patterns := strings.Split(lines[0], ", ")
	fmt.Println(Sum(Map(lines[2:], func(design string) int {
		return computeNumberOfTowelDesigns(design, patterns)
	})))
}

func computeNumberOfTowelDesigns(design string, patterns []string) int {
	count := make([]int, len(design)+1)
	count[0] = 1
	for i := 0; i < len(design); i++ {
		if count[i] == 0 {
			continue
		}
		for _, pattern := range patterns {
			if i+len(pattern) <= len(design) && design[i:i+len(pattern)] == pattern {
				count[i+len(pattern)] += count[i]
			}
		}
	}
	return count[len(design)]
}
