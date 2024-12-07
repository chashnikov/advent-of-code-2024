package main

import (
	"fmt"
	"regexp"
)

func day3(fileName string) {
	text := ReadText(fileName)
	mul := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	matches := mul.FindAllStringSubmatch(text, -1)
	res := 0
	for _, match := range matches {
		res += StrToInt(match[1]) * StrToInt(match[2])
	}
	fmt.Println(res)
}
