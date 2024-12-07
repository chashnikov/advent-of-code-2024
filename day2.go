package main

import (
	"fmt"
)

func day2(fileName string) {
	lines := ReadLines(fileName)
	safe := 0
	for _, line := range lines {
		nums := StrToIntegers(line)
		if IsSafe(nums) {
			safe++
		}
	}
	fmt.Println(safe)
}

func IsSafe(nums []int) bool {
	diffs := Diffs(nums)
	monotonic := ZipMap(diffs, diffs[1:], func(a int, b int) bool { return Sign(a) == Sign(b) })
	return All(monotonic, func(e bool) bool { return e }) &&
		All(diffs, func(d int) bool { return d != 0 && Abs(d) <= 3 })
}

func Diffs(nums []int) []int {
	return ZipMap(nums, nums[1:], func(a int, b int) int { return a - b })
}
