package main

import (
	"fmt"
	"strings"
)

func main() {
	lines := ReadLines("day1-full.txt")
	left := make([]int, 10)
	right := make([]int, 10)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		index := strings.Index(line, " ")
		leftNum := StrToInt(line[:index])
		rightNum := StrToInt(line[index:])
		left = append(left, leftNum)
		right = append(right, rightNum)
	}
	counts := make(map[int]int)
	for _, r := range right {
		counts[r]++
	}
	similarity := 0
	for _, l := range left {
		similarity += l * counts[l]
	}
	fmt.Println(similarity)
}
