package main

import (
	"slices"
	"strings"
)

func findCharInStrings(grid *[]string, char uint8) intVector {
	for y := 0; y < len(*grid); y++ {
		x := strings.IndexByte((*grid)[y], char)
		if x != -1 {
			return intVector{x, y}
		}
	}
	panic("Char not found: " + string(char))
}

func findCharInSlices(grid *[][]uint8, char uint8) intVector {
	for y := 0; y < len(*grid); y++ {
		x := slices.Index((*grid)[y], char)
		if x != -1 {
			return intVector{x, y}
		}
	}
	panic("Char not found: " + string(char))
}
