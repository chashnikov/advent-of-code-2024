package main

import "fmt"

func day10(fileName string) {
	lines := ReadLines(fileName)
	width := len(lines[0])
	height := len(lines)
	nines := make([]intVector, 0, 10)
	nineToIndex := make(map[intVector]uint)
	trailheads := make([][]*bitset, height)
	for y := 0; y < height; y++ {
		trailheads[y] = make([]*bitset, width)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			trailheads[y][x] = newBitset()
			if lines[y][x] == '9' {
				p := intVector{x, y}
				index := uint(len(nines))
				nineToIndex[p] = index
				nines = append(nines, p)
				trailheads[y][x].Set(index, true)
			}
		}
	}

	for digit := 8; digit >= 0; digit-- {
		char := '0' + byte(digit)
		prevChar := '0' + byte(digit+1)
		for x := 0; x < width; x++ {
			for y := 0; y < height; y++ {
				if lines[y][x] == char {
					if x > 0 && lines[y][x-1] == prevChar {
						trailheads[y][x].OrWith(trailheads[y][x-1])
					}
					if x < width-1 && lines[y][x+1] == prevChar {
						trailheads[y][x].OrWith(trailheads[y][x+1])
					}
					if y > 0 && lines[y-1][x] == prevChar {
						trailheads[y][x].OrWith(trailheads[y-1][x])
					}
					if y < height-1 && lines[y+1][x] == prevChar {
						trailheads[y][x].OrWith(trailheads[y+1][x])
					}
				}
			}
		}
	}

	score := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if lines[y][x] == '0' {
				score += trailheads[y][x].BitCount()
			}
		}
	}
	fmt.Println(score)
}
