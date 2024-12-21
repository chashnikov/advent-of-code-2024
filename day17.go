package main

import (
	"fmt"
	"strconv"
	"strings"
)

func day17(fileName string) {
	lines := ReadLines(fileName)
	a := StrToInt(strings.TrimPrefix(lines[0], "Register A: "))
	b := StrToInt(strings.TrimPrefix(lines[1], "Register B: "))
	c := StrToInt(strings.TrimPrefix(lines[2], "Register C: "))
	program := SplitToInts(strings.TrimPrefix(lines[4], "Program: "), ",")
	ip := 0
	output := make([]uint8, 0)
	for ip < len(program)-1 {
		var operand int
		switch program[ip+1] {
		case ip:
		case 0, 1, 2, 3:
			operand = program[ip+1]
		case 4:
			operand = a
		case 5:
			operand = b
		case 6:
			operand = c
		case 7:
			panic("Not Implemented")
		}
		switch program[ip] {
		case 0:
			a = a / (1 << operand)
		case 1:
			b = b ^ program[ip+1]
		case 2:
			b = operand & 7
		case 3:
			if a != 0 {
				ip = program[ip+1] - 2
			}
		case 4:
			b ^= c
		case 5:
			output = append(output, uint8(operand&7))
		case 6:
			b = a / (1 << operand)
		case 7:
			c = a / (1 << operand)
		}

		ip += 2
	}
	fmt.Println(strings.Join(Map(output, func(o uint8) string {
		return strconv.Itoa(int(o))
	}), ","))
}
