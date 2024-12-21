package main

import (
	"fmt"
	"regexp"
)

func day13(fileName string) {
	lines := ReadLines(fileName)
	buttonRegexp := regexp.MustCompile("Button .: X\\+(\\d+), Y\\+(\\d+)")
	prizeRegexp := regexp.MustCompile("Prize: X=(\\d+), Y=(\\d+)")
	res := int64(0)
	for m := 0; m < len(lines); m += 4 {
		aMatch := buttonRegexp.FindStringSubmatch(lines[m])
		bMatch := buttonRegexp.FindStringSubmatch(lines[m+1])
		prizeMatch := prizeRegexp.FindStringSubmatch(lines[m+2])
		aX := StrToInt64(aMatch[1])
		aY := StrToInt64(aMatch[2])
		bX := StrToInt64(bMatch[1])
		bY := StrToInt64(bMatch[2])
		pX := StrToInt64(prizeMatch[1]) + 10000000000000
		pY := StrToInt64(prizeMatch[2]) + 10000000000000
		res += ComputeMinimumTokens(aX, aY, bX, bY, pX, pY)
	}
	fmt.Println(res)
}

func ComputeMinimumTokens(aX int64, aY int64, bX int64, bY int64, pX int64, pY int64) int64 {
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
