package main

import (
	"fmt"
	"regexp"
)

func day13(fileName string) {
	lines := ReadLines(fileName)
	buttonRegexp := regexp.MustCompile("Button .: X\\+(\\d+), Y\\+(\\d+)")
	prizeRegexp := regexp.MustCompile("Prize: X=(\\d+), Y=(\\d+)")
	res := 0
	for m := 0; m < len(lines); m += 4 {
		aMatch := buttonRegexp.FindStringSubmatch(lines[m])
		bMatch := buttonRegexp.FindStringSubmatch(lines[m+1])
		prizeMatch := prizeRegexp.FindStringSubmatch(lines[m+2])
		aX := StrToInt(aMatch[1])
		aY := StrToInt(aMatch[2])
		bX := StrToInt(bMatch[1])
		bY := StrToInt(bMatch[2])
		pX := StrToInt(prizeMatch[1])
		pY := StrToInt(prizeMatch[2])
		res += ComputeMinimumTokens(aX, aY, bX, bY, pX, pY)
	}
	fmt.Println(res)
}

func ComputeMinimumTokens(aX int, aY int, bX int, bY int, pX int, pY int) int {
	// a*aX+b*bX = pX
	// a*aY+b*bY = pY
	det := aX*bY - aY*bX
	if det == 0 {
		return 0
	}
	a := (bY*pX - bX*pY) / det
	b := (-aY*pX + aX*pY) / det
	if a*aX+b*bX != pX || a*aY+b*bY != pY {
		return 0
	}
	return 3*a + b
}
