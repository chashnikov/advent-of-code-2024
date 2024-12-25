package main

import "fmt"

func day22(fileName string) {
	lines := ReadLines(fileName)
	res := Sum(Map(lines, func(line string) int {
		return generateSecretNumber(StrToInt(line), 2000)
	}))
	fmt.Println(res)
}

func generateSecretNumber(initial int, iterations int) int {
	secret := initial
	for i := 0; i < iterations; i++ {
		secret = (secret<<6 ^ secret) & 0xffffff
		secret = (secret>>5 ^ secret) & 0xffffff
		secret = (secret<<11 ^ secret) & 0xffffff
	}
	return secret
}
