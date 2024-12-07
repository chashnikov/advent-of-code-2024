package main

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(name string) []string {
	text := ReadText(name)
	if strings.Contains(text, "\r\n") {
		return strings.Split(text, "\r\n")
	}
	return strings.Split(text, "\n")
}

func ReadText(name string) string {
	data, err := os.ReadFile("inputs/" + name)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func SplitToInts(line string, separator string) []int {
	return Map(strings.Split(line, separator), StrToInt)
}

func StrToInt(s string) int {
	leftNum, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return leftNum
}
