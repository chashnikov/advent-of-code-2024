package main

import (
	"fmt"
	"slices"
	"strconv"
)

func day11(filePath string) {
	initial := SplitToInts(ReadText(filePath), " ")
	stones := Map(initial, func(i int) [2]int64 { return [2]int64{int64(i), 1} })

	for i := 1; i <= 75; i++ {
		stones = ReplaceStones(stones)
		fmt.Printf("%d: %d different stones\n", i, len(stones))
	}
	fmt.Println(Sum64(Map(stones, func(s [2]int64) int64 { return s[1] })))
}

func ReplaceStones(stones [][2]int64) [][2]int64 {
	newStones := make([][2]int64, 0)
	for _, stone := range stones {
		s := strconv.FormatInt(stone[0], 10)
		if stone[0] == 0 {
			newStones = append(newStones, [2]int64{1, stone[1]})
		} else if len(s)%2 == 0 {
			newStones = append(newStones, [2]int64{StrToInt64(s[:len(s)/2]), stone[1]})
			newStones = append(newStones, [2]int64{StrToInt64(s[len(s)/2:]), stone[1]})
		} else {
			newStones = append(newStones, [2]int64{2024 * stone[0], stone[1]})
		}
	}
	slices.SortFunc(newStones, func(a, b [2]int64) int {
		return int(a[0] - b[0])
	})
	merged := make([][2]int64, 0)
	for i := 0; i < len(newStones); i++ {
		if i == 0 || newStones[i][0] != newStones[i-1][0] {
			merged = append(merged, newStones[i])
		} else {
			merged[len(merged)-1][1] += newStones[i][1]
		}
	}
	return merged
}
