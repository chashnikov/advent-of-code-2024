package main

import (
	"fmt"
	"regexp"
)

func day14(fileName string) {
	lines := ReadLines(fileName)
	lineRegexp := regexp.MustCompile("p=(-?\\d+),(-?\\d+) v=(-?\\d+),(-?\\d+)")
	seconds := 100
	width := 101
	height := 103
	if fileName == "day14.txt" {
		width = 11
		height = 7
	}
	count := [2][2]int{}
	for _, line := range lines {
		matches := lineRegexp.FindStringSubmatch(line)
		p0 := intVector{StrToInt(matches[1]), StrToInt(matches[2])}
		v := intVector{StrToInt(matches[3]), StrToInt(matches[4])}
		p := p0.Plus(v.Multiply(seconds))
		x := (p.x%width + width) % width
		y := (p.y%height + height) % height
		if x == width/2 || y == height/2 {
			continue
		}
		ix := (x - 1) / (width / 2)
		iy := (y - 1) / (height / 2)
		count[ix][iy]++
	}
	fmt.Println(count[0][0] * count[0][1] * count[1][0] * count[1][1])
}
