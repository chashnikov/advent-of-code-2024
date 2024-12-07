package main

import "fmt"

func day4(fileName string) {
	lines := ReadLines(fileName)
	width := len(lines[0])
	height := len(lines)

	matches := 0
	for x := 1; x < width-1; x++ {
		for y := 1; y < height-1; y++ {
			if lines[y][x] != 'A' {
				continue
			}
			diag1 := string(lines[y-1][x-1]) + string(lines[y+1][x+1])
			diag2 := string(lines[y+1][x-1]) + string(lines[y-1][x+1])
			if (diag1 == "MS" || diag1 == "SM") && (diag2 == "MS" || diag2 == "SM") {
				matches++
			}
		}
	}
	fmt.Println(matches)
}
