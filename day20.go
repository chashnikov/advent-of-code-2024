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
	var minToSave int
	if fileName == "day20-full.txt" {
		minToSave = 100
	} else {
		minToSave = 50
	}
	cheats := 0
	cheatsCounts := make(map[int]int)
	processed := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == '#' {
				continue
			}
			cheat1 := intVector{x, y}
			cheats += computeNumberOfCheats(cheat1, grid, fromEnd, minToSave, cheatsCounts)
			processed++
		}
	}
	fmt.Println(cheats)
	/*	for saved := minToSave; saved < 1000; saved++ {
			if cheatsCounts[saved] > 0 {
				fmt.Printf("There are %d cheats that save %d picoseconds.\n", cheatsCounts[saved], saved)
			}
		}
	*/
}

type pointAndDist struct {
	point intVector
	dist  int
}

func computeNumberOfCheats(cheat1 intVector, grid []string, fromEnd [][]int, minToSave int, cheatsCounts map[int]int) int {
	height, width := len(grid), len(grid[0])
	oldToEnd := fromEnd[cheat1.y][cheat1.x]
	if oldToEnd == math.MaxInt {
		return 0
	}
	visited := CreateBool2dSlice(height, width)
	queue := make([]pointAndDist, 100)
	queue[0] = pointAndDist{cheat1, 0}
	queue = queue[:1]
	cheats := 0
	for len(queue) > 0 {
		p := queue[0].point
		dist := queue[0].dist
		queue = queue[1:]
		if dist < 20 {
			for _, dir := range directions4 {
				next := p.Plus(dir)
				if next.x < 0 || next.x >= width || next.y < 0 || next.y >= height || visited[next.y][next.x] {
					continue
				}
				visited[next.y][next.x] = true
				queue = append(queue, pointAndDist{next, dist + 1})
			}
		}
		if grid[p.y][p.x] != '#' {
			toEndNew := fromEnd[p.y][p.x]
			if toEndNew < math.MaxInt && toEndNew+dist < oldToEnd {
				saved := oldToEnd - (toEndNew + dist)
				if saved >= minToSave {
					if saved > 0 {
						cheatsCounts[saved]++
						//if saved == 72 {
						//	fmt.Printf("Cheat %d,%d -> %d,%d = %d\n", cheat1.x, cheat1.y, p.x, p.y, saved)
						//}
					}
					cheats++
				}
			}
		}
	}
	return cheats
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
