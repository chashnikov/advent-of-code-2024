package main

import "fmt"

type direction struct {
	dx int
	dy int
}

func day4(fileName string) {
	lines := ReadLines(fileName)
	width := len(lines[0])
	height := len(lines)

	allDirections := []direction{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	wordSearch := "XMAS"
	dLen := len(wordSearch) - 1
	matches := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			for _, d := range allDirections {
				if x+d.dx*dLen < 0 || x+d.dx*dLen >= width ||
					y+d.dy*dLen < 0 || y+d.dy*dLen >= height {
					continue
				}
				word := make([]byte, len(wordSearch))
				for i := 0; i <= dLen; i++ {
					word[i] = lines[y+d.dy*i][x+d.dx*i]
				}
				if wordSearch == string(word) {
					matches++
				}
			}
		}
	}
	fmt.Println(matches)
}
