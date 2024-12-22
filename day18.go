package main

import (
	"fmt"
	"math"
)

func day18(fileName string) {
	lines := ReadLines(fileName)
	points := Map(lines, func(line string) intVector {
		ints := SplitToInts(line, ",")
		return intVector{ints[0], ints[1]}
	})
	var size int
	var pointsLen int
	if fileName == "day18-full.txt" {
		size = 71
		pointsLen = 1024
	} else {
		size = 7
		pointsLen = 12
	}
	steps := make([][]int, size)
	grid := make([][]uint8, size)
	for y := 0; y < size; y++ {
		steps[y] = make([]int, size)
		grid[y] = make([]uint8, size)
		for x := 0; x < size; x++ {
			steps[y][x] = math.MaxInt
			grid[y][x] = '.'
		}
	}
	for i := 0; i < pointsLen; i++ {
		grid[points[i].y][points[i].x] = '#'
	}
	goInMemorySpace(intVector{0, 0}, 0, &grid, &steps)
	fmt.Println(steps[size-1][size-1])
}

func goInMemorySpace(p intVector, step int, grid *[][]uint8, steps *[][]int) {
	if p.x < 0 || p.x >= len((*steps)[0]) || p.y < 0 || p.y >= len(*steps) {
		return
	}
	if (*grid)[p.y][p.x] == '#' {
		return
	}
	if step >= (*steps)[p.y][p.x] {
		return
	}
	(*steps)[p.y][p.x] = step
	for _, direction := range directions4 {
		goInMemorySpace(p.Plus(direction), step+1, grid, steps)
	}
}
