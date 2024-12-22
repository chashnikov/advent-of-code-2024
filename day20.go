package main

import (
	"fmt"
	"math"
)

func day20(fileName string) {
	grid := ReadLines(fileName)
	width := len(grid[0])
	height := len(grid)
	start := findCharInStrings(&grid, 'S')
	end := findCharInStrings(&grid, 'E')
	fromStart := CreateInt2dSlice(height, width, math.MaxInt)
	fromEnd := CreateInt2dSlice(height, width, math.MaxInt)
	goAndComputeDistance(start, 0, &grid, &fromStart)
	goAndComputeDistance(end, 0, &grid, &fromEnd)
	cheats := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] != '#' {
				continue
			}
			cheat1 := intVector{x, y}
			minFromStart := math.MaxInt
			minDir := directions4[0]
			toEnd := math.MaxInt
			for _, dir := range directions4 {
				p := cheat1.Plus(dir)
				if p.x >= 0 && p.x < width && p.y >= 0 && p.y < height && fromStart[p.y][p.x] < minFromStart {
					minFromStart = fromStart[p.y][p.x]
					minDir = dir
					toEnd = fromEnd[p.y][p.x]
				}
			}
			if minFromStart == math.MaxInt {
				continue
			}
			for _, dir := range directions4 {
				if dir == minDir {
					continue
				}
				cheat2 := cheat1.Plus(dir)
				if cheat2.x >= 0 && cheat2.x < width && cheat2.y >= 0 && cheat2.y < height {
					toEndNew := fromEnd[cheat2.y][cheat2.x]
					if toEndNew < math.MaxInt && toEndNew+2 < toEnd {
						saved := toEnd - (toEndNew + 2)
						fmt.Printf("Cheat %d,%d->%d,%d save %d picoseconds\n", cheat1.x, cheat1.y, cheat2.x, cheat2.y, saved)
						if saved >= 100 {
							cheats++
						}
					}
				}
			}
		}
	}
	fmt.Println(cheats)
}

func goAndComputeDistance(p intVector, dist int, grid *[]string, distances *[][]int) {
	if p.x < 0 || p.x >= len((*grid)[0]) || p.y < 0 || p.y >= len(*grid) {
		return
	}
	if (*grid)[p.y][p.x] == '#' {
		return
	}
	if dist >= (*distances)[p.y][p.x] {
		return
	}
	(*distances)[p.y][p.x] = dist
	for _, dir := range directions4 {
		goAndComputeDistance(p.Plus(dir), dist+1, grid, distances)
	}
}
