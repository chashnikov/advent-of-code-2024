package main

import (
	"os"
	"strconv"
	"strings"
)

func ReadToString(name string) string {
	data, err := os.ReadFile("inputs/" + name)
	if err != nil {
		panic(err)
	}
	return string(data)
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
