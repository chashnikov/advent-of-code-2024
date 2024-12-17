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
	price := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			data := GoAndMarkRegion(x, y, &lines, lines[y][x], region, &regions)
			if data.area > 0 {
				region++
				price += data.area * data.perimeter
			}
		}
	}
	fmt.Println(price)
}

func GoAndMarkRegion(x int, y int, grid *[]string, plant uint8, region int, regions *[][]int) regionData {
	width := len((*grid)[0])
	height := len(*grid)
	if x < 0 || x >= width || y < 0 || y >= height || ((*regions)[y][x] != 0 && (*regions)[y][x] != region) || plant != (*grid)[y][x] {
		return regionData{0, 1}
	}
	if (*regions)[y][x] == region {
		return regionData{0, 0}
	}
	(*regions)[y][x] = region
	d1 := GoAndMarkRegion(x+1, y, grid, plant, region, regions)
	d2 := GoAndMarkRegion(x-1, y, grid, plant, region, regions)
	d3 := GoAndMarkRegion(x, y+1, grid, plant, region, regions)
	d4 := GoAndMarkRegion(x, y-1, grid, plant, region, regions)
	area := 1 + d1.area + d2.area + d3.area + d4.area
	perimeter := d1.perimeter + d2.perimeter + d3.perimeter + d4.perimeter
	return regionData{area, perimeter}
}

type regionData struct {
	area      int
	perimeter int
}
