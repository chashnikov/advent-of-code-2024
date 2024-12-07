package main

import (
	"os"
	"strconv"
	"strings"
)

func ReadLines(name string) []string {
	data, err := os.ReadFile("inputs/" + name)
	if err != nil {
		panic(err)
	}
	text := string(data)
	if strings.Contains(text, "\r\n") {
		return strings.Split(text, "\r\n")
	}
	return strings.Split(text, "\n")
}

func StrToIntegers(line string) []int {
	return Map(strings.Split(line, " "), StrToInt)
}

func StrToInt(s string) int {
	leftNum, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return leftNum
}
