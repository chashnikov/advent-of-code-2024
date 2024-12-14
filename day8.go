package main

import "fmt"

func day8(fileName string) {
	lines := ReadLines(fileName)
	antennas := make(map[uint8][]intVector)
	width := len(lines[0])
	height := len(lines)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			char := lines[y][x]
			if char != '.' {
				antennas[char] = append(antennas[char], intVector{x, y})
			}
		}
	}
	antinodes := make(map[intVector]bool)
	for _, points := range antennas {
		for i := 0; i < len(points); i++ {
			for j := i + 1; j < len(points); j++ {
				ij := points[j].Minus(points[i])
				a1 := points[i].Minus(ij)
				if a1.x >= 0 && a1.x < width && a1.y >= 0 && a1.y < height {
					antinodes[a1] = true
				}
				a2 := points[j].Plus(ij)
				if a2.x >= 0 && a2.x < width && a2.y >= 0 && a2.y < height {
					antinodes[a2] = true
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
