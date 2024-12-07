package main

import "os"

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
