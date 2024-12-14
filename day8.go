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
				ijMin := ij.Divide(Gcd(ij.x, ij.y))
				p := points[i]
				for {
					if p.x >= 0 && p.x < width && p.y >= 0 && p.y < height {
						antinodes[p] = true
					} else {
						break
					}
					p = p.Plus(ijMin)
				}
				p = points[i].Minus(ijMin)
				for {
					if p.x >= 0 && p.x < width && p.y >= 0 && p.y < height {
						antinodes[p] = true
					} else {
						break
					}
					p = p.Minus(ijMin)
				}
			}
		}
	}
	fmt.Println(len(antinodes))
}
