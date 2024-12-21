package main

import (
	"fmt"
	"regexp"
	"time"
)

func day14(fileName string) {
	lines := ReadLines(fileName)
	lineRegexp := regexp.MustCompile("p=(-?\\d+),(-?\\d+) v=(-?\\d+),(-?\\d+)")
	width := 101
	height := 103
	if fileName == "day14.txt" {
		width = 11
		height = 7
	}
	p0 := make([]intVector, 0)
	v := make([]intVector, 0)
	for _, line := range lines {
		matches := lineRegexp.FindStringSubmatch(line)
		p0i := intVector{StrToInt(matches[1]), StrToInt(matches[2])}
		vi := intVector{StrToInt(matches[3]), StrToInt(matches[4])}
		p0 = append(p0, p0i)
		v = append(v, vi)
	}
	grid := make([][]int, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]int, width)
	}
	for second := 0; ; second++ {
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				grid[y][x] = 0
			}
		}
		for r := 0; r < len(p0); r++ {
			p := p0[r].Plus(v[r].Multiply(second))
			x := (p.x%width + width) % width
			y := (p.y%height + height) % height
			grid[y][x]++
		}
		if !LooksLikeATree(&grid) {
			continue
		}
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if grid[y][x] > 0 {
					fmt.Print("X")
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println()
		}
		fmt.Println()
		fmt.Println(second)
		time.Sleep(1 * time.Second)
	}
}

func LooksLikeATree(grid *[][]int) bool {
	height := len(*grid)
	width := len((*grid)[0])
	regions := make([][]int, height)
	for y := 0; y < height; y++ {
		regions[y] = make([]int, width)
	}
	region := 1
	for y := 0; y < height-1; y++ {
		for x := 0; x < width-1; x++ {
			area := ComputeGridArea(x, y, grid, region, &regions)
			if area > 120 {
				return true
			}
			if area > 0 {
				region++
			}
		}
	}
	return false
}

func ComputeGridArea(x int, y int, grid *[][]int, region int, regions *[][]int) int {
	width := len((*grid)[0])
	height := len(*grid)
	if x < 0 || x >= width || y < 0 || y >= height || (*grid)[y][x] == 0 || (*regions)[y][x] == region {
		return 0
	}
	(*regions)[y][x] = region
	area1 := ComputeGridArea(x+1, y, grid, region, regions)
	area2 := ComputeGridArea(x-1, y, grid, region, regions)
	area3 := ComputeGridArea(x, y+1, grid, region, regions)
	area4 := ComputeGridArea(x, y-1, grid, region, regions)
	return 1 + area1 + area2 + area3 + area4
}
