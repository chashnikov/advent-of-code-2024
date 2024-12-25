package main

import "fmt"

func day22(fileName string) {
	lines := ReadLines(fileName)
	prices := Map(lines, func(line string) []int {
		return generatePricesList(StrToInt(line), 2000)
	})
	maxSum := 0
	for c0 := -9; c0 <= 9; c0++ {
		fmt.Printf("c0=%d\n", c0)
		for c1 := -9; c1 <= 9; c1++ {
			fmt.Printf("c1=%d\n", c1)
			for c2 := -9; c2 <= 9; c2++ {
				for c3 := -9; c3 <= 9; c3++ {
					sum := 0
					for i := 0; i < len(prices); i++ {
						sum += findFirstPrice(prices[i], c0, c1, c2, c3)
					}
					maxSum = max(maxSum, sum)
				}
			}
		}
	}
	fmt.Println(maxSum)
}

func findFirstPrice(prices []int, c0 int, c1 int, c2 int, c3 int) int {
	for i := 4; i < len(prices); i++ {
		if prices[i-3]-prices[i-4] == c0 && prices[i-2]-prices[i-3] == c1 &&
			prices[i-1]-prices[i-2] == c2 && prices[i]-prices[i-1] == c3 {
			return prices[i]
		}
	}
	return 0
}

func generatePricesList(initial int, iterations int) []int {
	secret := initial
	result := make([]int, iterations+1)
	result[0] = secret % 10
	for i := 1; i <= iterations; i++ {
		secret = (secret<<6 ^ secret) & 0xffffff
		secret = (secret>>5 ^ secret) & 0xffffff
		secret = (secret<<11 ^ secret) & 0xffffff
		result[i] = secret % 10
	}
	return result
}
