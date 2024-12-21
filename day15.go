package main

import (
	"fmt"
	"slices"
	"strings"
)

func day15(fileName string) {
	lines := ReadLines(fileName)
	grid := make([][]uint8, 0)
	y0 := -1
	x0 := -1
	height := slices.Index(lines, "")
	for y := 0; y < height; y++ {
		line := lines[y]
		position := strings.Index(line, "@")
		if position >= 0 {
			x0 = position
			y0 = len(grid)
		}
		chars := []uint8(line)
		grid = append(grid, chars)
	}
	position := intVector{x0, y0}
	for i := height + 1; i < len(lines); i++ {
		movements := lines[i]
		for _, c := range movements {
			dir := intVector{0, 0}
			switch c {
			case '>':
				dir = intVector{1, 0}
			case '<':
				dir = intVector{-1, 0}
			case '^':
				dir = intVector{0, -1}
			case 'v':
				dir = intVector{0, 1}
			default:
				panic("invalid movement: " + string(c))
			}
			if MoveRobotOrBox(position, dir, &grid) {
				position = position.Plus(dir)
			}
		}
	}
	width := len(grid[0])
	coordinates := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 'O' {
				coordinates += 100*y + x
			}
		}
	}
	fmt.Println(coordinates)
}

func MoveRobotOrBox(position intVector, dir intVector, grid *[][]uint8) bool {
	next := position.Plus(dir)
	item := (*grid)[next.y][next.x]
	if item == '#' {
		return false
	}
	if item == 'O' {
		if !MoveRobotOrBox(next, dir, grid) {
			return false
		}
	}
	(*grid)[next.y][next.x] = (*grid)[position.y][position.x]
	(*grid)[position.y][position.x] = '.'
	return true
}
