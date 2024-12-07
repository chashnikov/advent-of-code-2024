package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	text := ReadToString("day1.txt")
	lines := strings.Split(text, "\n")
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
	slices.Sort(left)
	slices.Sort(right)
	diff := 0
	for i, l := range left {
		diff += Abs(l - right[i])
	}
	fmt.Println(diff)
}
