package main

import (
	"fmt"
	"strings"
)

func day2(fileName string) {
	lines := ReadLines(fileName)
	safe := 0
loop:
	for _, line := range lines {
		nums := strings.Split(line, " ")
		var diff int
		var prev int
		var prevDiff int
		for i, numString := range nums {
			num := StrToInt(numString)
			if i > 0 {
				diff = num - prev
				if diff == 0 || Abs(diff) > 3 {
					continue loop
				}
				if i > 1 && Sign(diff) != Sign(prevDiff) {
					continue loop
				}
				prevDiff = diff
			}
			prev = num
		}
		safe++
	}
	fmt.Println(safe)
}
