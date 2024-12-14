package main

import "fmt"

func day9(fileName string) {
	text := ReadText(fileName)
	files := make([]*diskRange, 0)
	free := make([]*diskRange, 0)
	index := 0
	fileId := 0
	fill := true
	for _, char := range text {
		digit := int(char - '0')
		if fill {
			files = append(files, &diskRange{index, digit})
			fileId++
		} else {
			free = append(free, &diskRange{index, digit})
		}
		index += digit
		fill = !fill
	}
	fileId--
	for fileId > 0 {
		length := files[fileId].len
		for _, r := range free {
			if r.len >= length && r.start < files[fileId].start {
				files[fileId].start = r.start
				r.start += length
				r.len -= length
				break
			}
		}
		fileId--
	}
	res := int64(0)
	for id, file := range files {
		for i := 0; i < file.len; i++ {
			res += int64(file.start+i) * int64(id)
		}
	}
	fmt.Println(res)
}

type diskRange struct {
	start int
	len   int
}
