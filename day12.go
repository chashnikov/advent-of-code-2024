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
	edges := make(map[regionEdge]int)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if regions[y][x] != 0 {
				continue
			}
			p := intVector{x: x, y: y}
			area := GoAndMarkRegion(p, p, &lines, lines[y][x], region, &regions, &edges)
			if area > 0 {
				areas = append(areas, area)
				region++
			}
		}
	}

	price := 0
	for x := 0; x <= width; x++ {
		start := intVector{x, 0}
		direction := intVector{0, 1}
		ComputePrice(start, direction, height, &edges, &areas, &price)
		start = intVector{x, height}
		direction = intVector{0, -1}
		ComputePrice(start, direction, height, &edges, &areas, &price)
	}
	for y := 0; y <= height; y++ {
		start := intVector{0, y}
		direction := intVector{1, 0}
		ComputePrice(start, direction, width, &edges, &areas, &price)
		start = intVector{width, y}
		direction = intVector{-1, 0}
		ComputePrice(start, direction, height, &edges, &areas, &price)
	}
	fmt.Println(price)
}

func ComputePrice(start intVector, direction intVector, totalLen int, edges *map[regionEdge]int, areas *[]int, price *int) {
	prevRegion := -1
	for i := 0; i < totalLen; i++ {
		edge := regionEdge{start.Plus(direction.Multiply(i)), start.Plus(direction.Multiply(i + 1))}
		region, ok := (*edges)[edge]
		if !ok {
			prevRegion = -1
			continue
		}
		if prevRegion != region {
			*price += (*areas)[region]
		}
		prevRegion = region
	}
}

func GoAndMarkRegion(from intVector, to intVector, grid *[]string, plant uint8, region int, regions *[][]int, edges *map[regionEdge]int) int {
	width := len((*grid)[0])
	height := len(*grid)
	x := to.x
	y := to.y
	if x < 0 || x >= width || y < 0 || y >= height || plant != (*grid)[y][x] {
		edgeDirection := intVector{x: from.y - to.y, y: to.x - from.x}
		edgeStart := from.Plus(to).Plus(intVector{1, 1}).Minus(edgeDirection).Divide(2)
		edge := regionEdge{edgeStart, edgeStart.Plus(edgeDirection)}
		(*edges)[edge] = region
		return 0
	}
	if (*regions)[y][x] == region {
		return 0
	}
	(*regions)[y][x] = region
	area1 := GoAndMarkRegion(to, intVector{x + 1, y}, grid, plant, region, regions, edges)
	area2 := GoAndMarkRegion(to, intVector{x - 1, y}, grid, plant, region, regions, edges)
	area3 := GoAndMarkRegion(to, intVector{x, y + 1}, grid, plant, region, regions, edges)
	area4 := GoAndMarkRegion(to, intVector{x, y - 1}, grid, plant, region, regions, edges)
	return 1 + area1 + area2 + area3 + area4
}

type regionEdge struct {
	start intVector
	end   intVector
}
