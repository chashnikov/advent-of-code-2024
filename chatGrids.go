package main

import "strings"

func findChar(grid *[]string, char uint8) intVector {
	for y := 0; y < len(*grid); y++ {
		x := strings.IndexByte((*grid)[y], char)
		if x != -1 {
			return intVector{x, y}
		}
	}
	panic("Char not found: " + string(char))
}
