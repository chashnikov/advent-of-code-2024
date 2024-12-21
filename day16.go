package main

import (
	"fmt"
	"math"
	"slices"
)

func day16(fileName string) {
	grid := ReadLines(fileName)
	start := findChar(&grid, 'S')
	end := findChar(&grid, 'E')
	height := len(grid)
	width := len(grid[0])
	score := make([][][4]int, height)
	for y := 0; y < height; y++ {
		score[y] = make([][4]int, width)
		for x := 0; x < width; x++ {
			for d := 0; d < 4; d++ {
				score[y][x][d] = math.MaxInt
			}
		}
	}
	goReindeerMaze(start, 0, 0, &grid, &score)
	endScore := score[end.y][end.x]
	fmt.Println(slices.Min(endScore[:]))
}

func goReindeerMaze(p intVector, d int, currentScore int, grid *[]string, score *[][][4]int) {
	if p.x < 0 || p.x >= len((*score)[0]) || p.y < 0 || p.y >= len(*score) {
		return
	}
	if (*score)[p.y][p.x][d] < currentScore {
		return
	}
	if (*grid)[p.y][p.x] == '#' {
		return
	}
	(*score)[p.y][p.x][d] = currentScore
	goReindeerMaze(p, (d+1)%4, currentScore+1000, grid, score)
	goReindeerMaze(p, (d+3)%4, currentScore+1000, grid, score)
	goReindeerMaze(p.Plus(directions4[d]), d, currentScore+1, grid, score)
}
