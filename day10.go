package main

import "fmt"

func day10(fileName string) {
	lines := ReadLines(fileName)
	width := len(lines[0])
	height := len(lines)
	ratings := make([][]uint64, height)
	for y := 0; y < height; y++ {
		ratings[y] = make([]uint64, width)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			ratings[y][x] = 0
			if lines[y][x] == '9' {
				ratings[y][x] = 1
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
						ratings[y][x] += ratings[y][x-1]
					}
					if x < width-1 && lines[y][x+1] == prevChar {
						ratings[y][x] += ratings[y][x+1]
					}
					if y > 0 && lines[y-1][x] == prevChar {
						ratings[y][x] += ratings[y-1][x]
					}
					if y < height-1 && lines[y+1][x] == prevChar {
						ratings[y][x] += ratings[y+1][x]
					}
				}
			}
		}
	}

	var score uint64 = 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if lines[y][x] == '0' {
				score += ratings[y][x]
			}
		}
	}
	fmt.Println(score)
}
