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
	if fileName == "day18-full.txt" {
		size = 71
	} else {
		size = 7
	}
	steps := make([][]int, size)
	grid := make([][]uint8, size)
	for y := 0; y < size; y++ {
		steps[y] = make([]int, size)
		grid[y] = make([]uint8, size)
	}
	from := 0
	to := len(points)
	for {
		for y := 0; y < size; y++ {
			for x := 0; x < size; x++ {
				steps[y][x] = math.MaxInt
				grid[y][x] = '.'
			}
		}
		mid := (from + to) / 2
		for i := 0; i < mid; i++ {
			grid[points[i].y][points[i].x] = '#'
		}
		goInMemorySpace(intVector{0, 0}, 0, &grid, &steps)
		path := steps[size-1][size-1]
		if path == math.MaxInt {
			fmt.Printf("Not reachable, mid=%d\n", mid)
			if from == mid-1 {
				fmt.Printf("%d,%d\n", points[mid-1].x, points[mid-1].y)
				break
			}
			to = mid
		} else {
			fmt.Printf("Reachable, mid=%d\n", mid)
			if to == mid+1 {
				fmt.Printf("%d,%d\n", points[mid].x, points[mid].y)
				break
			}
			from = mid
		}
	}
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
