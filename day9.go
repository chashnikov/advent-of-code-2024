package main

import "fmt"

func day9(fileName string) {
	text := ReadText(fileName)
	disk := make([]int, 10*len(text))
	index := 0
	fileId := 0
	end := 0
	fill := true
	for _, char := range text {
		digit := int(char - '0')
		if fill {
			for i := 0; i < digit; i++ {
				disk[index+i] = fileId
			}
			end = index + digit - 1
			fileId++
		} else {
			for i := 0; i < digit; i++ {
				disk[index+i] = -1
			}
		}
		index += digit
		fill = !fill
	}
	res := 0
	index = 0
	for index <= end {
		if disk[index] == -1 {
			for end >= index && disk[end] == -1 {
				end--
			}
			if index > end {
				break
			}
			disk[index] = disk[end]
			end--
		}
		res += disk[index] * index
		index++
	}
	fmt.Println(res)
}
