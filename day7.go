package main

import (
	"fmt"
	"strings"
)

func day7(fileName string) {
	lines := ReadLines(fileName)
	var result int64 = 0
	for _, line := range lines {
		colon := strings.Index(line, ":")
		value := StrToInt64(line[0:colon])
		operands := Map(strings.Split(strings.TrimSpace(line[colon+1:]), " "), func(s string) int64 {
			return StrToInt64(strings.TrimSpace(s))
		})
		if CanBeComposedFrom(value, operands) {
			result += value
		}
	}
	fmt.Println(result)
}

func CanBeComposedFrom(value int64, operands []int64) bool {
	if len(operands) == 1 {
		return value == operands[0]
	}
	last := operands[len(operands)-1]
	rest := operands[:len(operands)-1]
	if value > last && CanBeComposedFrom(value-last, rest) {
		return true
	}
	if last > 0 && value%last == 0 && CanBeComposedFrom(value/last, rest) {
		return true
	}
	return false
}
