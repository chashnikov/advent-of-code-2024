package main

import (
	"fmt"
	"math"
	"slices"
)

func day16(fileName string) {
	grid := Map(ReadLines(fileName), func(s string) []uint8 {
		return []uint8(s)
	})
	start := findCharInSlices(&grid, 'S')
	end := findCharInSlices(&grid, 'E')
	height := len(grid)
	width := len(grid[0])
	score := make([][][4]int, height)
	prev := make([][][4][]uint8, height)
	for y := 0; y < height; y++ {
		score[y] = make([][4]int, width)
		prev[y] = make([][4][]uint8, width)
		for x := 0; x < width; x++ {
			for d := 0; d < 4; d++ {
				score[y][x][d] = math.MaxInt
				prev[y][x][d] = make([]uint8, 0)
			}
		}
	}
	goReindeerMaze(start, 0, 0, 0, &grid, &score, &prev)
	endScore := score[end.y][end.x]
	minScore := slices.Min(endScore[:])
	for d := 0; d < 4; d++ {
		if endScore[d] == minScore {
			goBackInReindeerMaze(end, d, &grid, &score, &prev)
		}
	}
	fmt.Println(Sum(Map(grid, func(row []uint8) int {
		return Count(row, func(char uint8) bool {
			return char == 'O'
		})
	})))
}

func goBackInReindeerMaze(p intVector, d int, grid *[][]uint8, score *[][][4]int, prev *[][][4][]uint8) {
	if (*score)[p.y][p.x][d] == -1 {
		return
	}
	(*score)[p.y][p.x][d] = -1
	(*grid)[p.y][p.x] = 'O'
	moves := (*prev)[p.y][p.x][d]
	for _, move := range moves {
		switch move {
		case 1:
			goBackInReindeerMaze(p, (d+3)%4, grid, score, prev)
		case 2:
			goBackInReindeerMaze(p, (d+1)%4, grid, score, prev)
		case 3:
			goBackInReindeerMaze(p.Minus(directions4[d]), d, grid, score, prev)
		}
	}
}

func goReindeerMaze(p intVector, d int, currentScore int, move uint8, grid *[][]uint8, score *[][][4]int, prev *[][][4][]uint8) {
	if p.x < 0 || p.x >= len((*score)[0]) || p.y < 0 || p.y >= len(*score) {
		return
	}
	if (*grid)[p.y][p.x] == '#' {
		return
	}
	oldScore := (*score)[p.y][p.x][d]
	if oldScore < currentScore {
		return
	}
	if oldScore == currentScore {
		(*prev)[p.y][p.x][d] = append((*prev)[p.y][p.x][d], move)
		return
	}
	(*prev)[p.y][p.x][d] = []uint8{move}
	(*score)[p.y][p.x][d] = currentScore
	goReindeerMaze(p, (d+1)%4, currentScore+1000, 1, grid, score, prev)
	goReindeerMaze(p, (d+3)%4, currentScore+1000, 2, grid, score, prev)
	goReindeerMaze(p.Plus(directions4[d]), d, currentScore+1, 3, grid, score, prev)
}
