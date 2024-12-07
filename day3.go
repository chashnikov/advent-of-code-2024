package main

import (
	"fmt"
	"regexp"
	"strings"
)

func day3(fileName string) {
	text := ReadText(fileName)
	mul := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")
	res := 0
	for {
		match := mul.FindStringSubmatch(text)
		if match == nil {
			break
		}
		nextIndex := strings.Index(text, match[0])
		disableIndex := strings.Index(text, "don't()")
		if disableIndex != -1 && disableIndex < nextIndex {
			text = text[disableIndex+1:]
			enableIndex := strings.Index(text, "do()")
			if enableIndex == -1 {
				break
			}
			text = text[enableIndex+1:]
			continue
		}
		res += StrToInt(match[1]) * StrToInt(match[2])
		text = text[nextIndex+1:]
	}
	fmt.Println(res)
}
