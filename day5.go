package main

import (
	"fmt"
	"strings"
)

func day5(fileName string) {
	lines := ReadLines(fileName)
	parsingOrdering := true
	orderings := make(map[[2]int]bool, 0)
	updates := make([][]int, 0)
	for index, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			parsingOrdering = false
			continue
		}
		if parsingOrdering {
			pair := SplitToInts(line, "|")
			if len(pair) != 2 {
				panic(index)
			}
			orderings[[2]int{pair[0], pair[1]}] = true
		} else {
			updates = append(updates, SplitToInts(line, ","))
		}
	}
	res := 0
	for _, update := range updates {
		if IsInRightOrder(update, orderings) {
			if len(update)%2 == 0 {
				panic(update)
			}
			res += update[(len(update)-1)/2]
		}
	}
	fmt.Println(res)
}

func IsInRightOrder(update []int, orderings map[[2]int]bool) bool {
	for i := 0; i < len(update); i++ {
		for j := i + 1; j < len(update); j++ {
			if orderings[[2]int{update[j], update[i]}] {
				return false
			}
		}
	}
	return true
}
