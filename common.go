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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func StrToInt(s string) int {
	leftNum, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return leftNum
}
