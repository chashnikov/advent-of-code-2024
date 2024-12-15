package main

import (
	"fmt"
	"strconv"
)

func day11(filePath string) {
	initial := SplitToInts(ReadText(filePath), " ")
	stones := Map(initial, func(i int) int64 { return int64(i) })
	for i := 1; i <= 25; i++ {
		stones = ReplaceStones(stones)
	}
	fmt.Println(len(stones))
}

func ReplaceStones(stones []int64) []int64 {
	newStones := make([]int64, 0)
	for _, stone := range stones {
		s := strconv.FormatInt(stone, 10)
		if stone == 0 {
			newStones = append(newStones, 1)
		} else if len(s)%2 == 0 {
			newStones = append(newStones, StrToInt64(s[:len(s)/2]))
			newStones = append(newStones, StrToInt64(s[len(s)/2:]))
		} else {
			newStones = append(newStones, 2024*stone)
		}
	}
	return newStones
}
