package main

import (
	"fmt"
	"slices"
)

func day6(fileName string) {
	grid := Map(ReadLines(fileName), func(s string) []byte {
		result := make([]byte, len(s))
		for i, c := range s {
			result[i] = byte(c)
		}
		return result
	})
	width := len(grid[0])
	height := len(grid)

	isInitial := func(c byte) bool { return c == '^' }
	y := slices.IndexFunc(grid, func(s []byte) bool {
		return Any(s, isInitial)
	})
	x := slices.IndexFunc(grid[y], isInitial)
	dirIndex := 0
	for {
		grid[y][x] = 'X'
		nx := x + directions[dirIndex].dx
		ny := y + directions[dirIndex].dy
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			break
		}
		if grid[ny][nx] == '#' {
			dirIndex = (dirIndex + 1) % 4
			continue
		}
		x = nx
		y = ny
	}
	result := Sum(Map(grid, func(row []byte) int {
		return Count(row, func(b byte) bool {
			return b == 'X'
		})
	}))
	fmt.Println(result)
}

var directions = []direction{
	{0, -1},
	{1, 0},
	{0, 1},
	{-1, 0},
}

type direction struct {
	dx int
	dy int
}
