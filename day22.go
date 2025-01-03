package main

import (
	"fmt"
	"runtime"
	"time"
)

func day22(fileName string) {
	start := time.Now()
	lines := ReadLines(fileName)
	iterations := 2000
	prices := Map(lines, func(line string) []int {
		return generatePricesList(StrToInt(line), iterations)
	})
	differences := Map(prices, func(values []int) []int {
		diff := make([]int, len(values)-1)
		for i := 0; i < len(values)-1; i++ {
			diff[i] = values[i+1] - values[i]
		}
		return diff
	})
	firstIndexes := Map(differences, func(diffs []int) [19][19]int {
		firstIndex := [19][19]int{}
		for i := 0; i < 19; i++ {
			for j := 0; j < 19; j++ {
				firstIndex[i][j] = len(diffs)
			}
		}
		for i := 0; i < len(diffs)-1; i++ {
			if firstIndex[diffs[i]+9][diffs[i+1]+9] == len(diffs) {
				firstIndex[diffs[i]+9][diffs[i+1]+9] = i
			}
		}
		return firstIndex
	})
	c := [4]int{-9, -9, -9, -9}
	inputs := make(chan [4]int, runtime.NumCPU())
	sums := make(chan int)
	inputsNum := 0
	for {
		inputsNum++
		inputs <- c
		go func() {
			changes := <-inputs
			sums <- computeSumPrices(prices, firstIndexes, changes, differences)
		}()

		i := 3
		for i >= 0 && c[i] == 9 {
			c[i] = -9
			i--
		}
		if i < 0 {
			break
		}
		if i == 1 {
			fmt.Printf("c = {%d, %d, }\n", c[0], c[1])
		}
		c[i]++
	}
	maxSum := 0
	for i := 0; i < inputsNum; i++ {
		sum := <-sums
		maxSum = max(maxSum, sum)
	}
	fmt.Printf("Sum for %d iterations: %d\n", iterations, maxSum)
	fmt.Printf("Elapsed time: %s\n", time.Since(start))
}

func computeSumPrices(prices [][]int, firstIndexes [][19][19]int, c [4]int, differences [][]int) int {
	sum := 0
	for i := 0; i < len(prices); i++ {
		if firstIndexes[i][c[1]+9][c[2]+9] == len(differences[i]) ||
			firstIndexes[i][c[2]+9][c[3]+9] == len(differences[i]) {
			continue
		}
		firstIndex := firstIndexes[i][c[0]+9][c[1]+9]
		if firstIndex < len(differences[i]) {
			sum += findFirstPrice(differences[i], prices[i], firstIndex, c)
		}
	}
	return sum
}

func findFirstPrice(differences []int, prices []int, firstIndex int, c [4]int) int {
	for i := firstIndex; i < len(differences)-3; i++ {
		if differences[i] == c[0] && differences[i+1] == c[1] &&
			differences[i+2] == c[2] && differences[i+3] == c[3] {
			return prices[i+4]
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
