package main

import (
	"fmt"
	"slices"
)

func day6(fileName string) {
	grid0 := Map(ReadLines(fileName), func(s string) []byte {
		result := make([]byte, len(s))
		for i, c := range s {
			result[i] = byte(c)
		}
		return result
	})
	width := len(grid0[0])
	height := len(grid0)

	isInitial := func(c byte) bool { return c == '^' }
	y0 := slices.IndexFunc(grid0, func(s []byte) bool {
		return Any(s, isInitial)
	})
	x0 := slices.IndexFunc(grid0[y0], isInitial)
	grid := Map(grid0, slices.Clone)
	simulateGuard(grid, x0, y0, width, height)

	result := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if grid[y][x] != 'X' || x == x0 && y == y0 {
				continue
			}
			newGrid := Map(grid0, slices.Clone)
			newGrid[y][x] = '#'
			if simulateGuard(newGrid, x0, y0, width, height) {
				result++
			}
		}
	}

	fmt.Println(result)
}

func simulateGuard(grid [][]byte, x int, y int, width int, height int) bool {
	dirIndex := 0
	moves := 0
	for {
		moves++
		if moves > width*height*4 {
			return true
		}
		grid[y][x] = 'X'
		nx := x + directions[dirIndex].dx
		ny := y + directions[dirIndex].dy
		if nx < 0 || nx >= width || ny < 0 || ny >= height {
			return false
		}
		if grid[ny][nx] == '#' {
			dirIndex = (dirIndex + 1) % 4
			continue
		}
		x = nx
		y = ny
	}
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
