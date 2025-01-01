package main

import "fmt"

func day25(fileName string) {
	lines := ReadLines(fileName)
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	for i := 0; i < len(lines); i += 8 {
		heights := [5]int{-1, -1, -1, -1, -1}
		for j := 0; j < 7; j++ {
			for k := 0; k < 5; k++ {
				if lines[i+j][k] == '#' {
					heights[k]++
				}
			}
		}
		if lines[i][0] == '#' {
			locks = append(locks, heights)
		} else {
			keys = append(keys, heights)
		}
	}
	result := 0
	for _, lock := range locks {
		for _, key := range keys {
			fit := true
			for k := 0; k < 5; k++ {
				if lock[k]+key[k] > 5 {
					fit = false
					break
				}
			}
			if fit {
				result++
			}
		}
	}
	fmt.Println(result)
}
