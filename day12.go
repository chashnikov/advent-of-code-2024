package main

import "fmt"

func day12(fileName string) {
	lines := ReadLines(fileName)
	width := len(lines[0])
	height := len(lines)
	regions := make([][]int, height)
	for y := 0; y < height; y++ {
		regions[y] = make([]int, width)
	}
	region := 1
	areas := make([]int, 1)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			area := GoAndMarkRegion(x, y, &lines, lines[y][x], region, &regions)
			if area > 0 {
				areas = append(areas, area)
				region++
			}
		}
	}
	sides := make(map[int]int)
	price := 0
	fmt.Println(price)
}

func GoAndMarkRegion(x int, y int, grid *[]string, plant uint8, region int, regions *[][]int) int {
	width := len((*grid)[0])
	height := len(*grid)
	if x < 0 || x >= width || y < 0 || y >= height || (*regions)[y][x] != 0 || plant != (*grid)[y][x] {
		return 0
	}
	(*regions)[y][x] = region
	area1 := GoAndMarkRegion(x+1, y, grid, plant, region, regions)
	area2 := GoAndMarkRegion(x-1, y, grid, plant, region, regions)
	area3 := GoAndMarkRegion(x, y+1, grid, plant, region, regions)
	area4 := GoAndMarkRegion(x, y-1, grid, plant, region, regions)
	return 1 + area1 + area2 + area3 + area4
}

type regionData struct {
	area      int
	perimeter int
}
