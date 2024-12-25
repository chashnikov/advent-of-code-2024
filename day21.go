package main

import (
	"fmt"
	"slices"
	"strings"
)

func day21(fileName string) {
	seq5 := "<v<A>>^AvA^A<vA<AA>>^AAvA<^A>AAvA^A<vA>^AA<A>A<v<A>A>^AAAvA<^A>A"
	fmt.Println(seq5)
	seq5 = typeKeySequence(seq5)
	fmt.Println(seq5)
	seq5 = typeKeySequence(seq5)
	fmt.Println(seq5)
	lines := ReadLines(fileName)
	fmt.Println(Sum(Map(lines, func(line string) int {
		num := StrToInt(strings.TrimSuffix(line, "A"))
		keySequence := ComputeKeySequence([]uint8(line), 0)
		fmt.Print(line)
		fmt.Print(":")
		fmt.Println(string(keySequence))
		return num * len(keySequence)
	})))
}

var keyCoordinates = map[uint8]intVector{
	'A': {0, 0},
	'0': {-1, 0},
	'1': {-2, -1},
	'2': {-1, -1},
	'3': {0, -1},
	'4': {-2, -2},
	'5': {-1, -2},
	'6': {0, -2},
	'7': {-2, -3},
	'8': {-1, -3},
	'9': {0, -3},
	'^': {-1, 0},
	'<': {-2, 1},
	'v': {-1, 1},
	'>': {0, 1},
}
var forbiddenCoordinate = intVector{-2, 0}

func ComputeKeySequence(line []uint8, device int) []uint8 {
	if device == 3 {
		return line
	}
	p := intVector{0, 0}
	seq := make([]uint8, 0)
	for _, ch := range line {
		next, ok := keyCoordinates[ch]
		if !ok {
			panic(fmt.Sprintf("invalid character '%c'", ch))
		}
		dx := next.x - p.x
		dy := next.y - p.y
		var horizontal []uint8
		if dx > 0 {
			horizontal = CreateSlice[uint8](dx, '>')
		} else {
			horizontal = CreateSlice[uint8](-dx, '<')
		}
		var vertical []uint8
		if dy > 0 {
			vertical = CreateSlice[uint8](dy, 'v')
		} else {
			vertical = CreateSlice[uint8](-dy, '^')
		}
		keySeq := make([]uint8, 0)
		if p.Plus(intVector{dx, 0}) != forbiddenCoordinate {
			keySeq = ComputeKeySequence(slices.Concat(horizontal, vertical, []uint8{'A'}), device+1)
		}
		if p.Plus(intVector{0, dy}) != forbiddenCoordinate {
			newKeySeq := ComputeKeySequence(slices.Concat(vertical, horizontal, []uint8{'A'}), device+1)
			if len(keySeq) == 0 || len(keySeq) > len(newKeySeq) {
				keySeq = newKeySeq
			}
		}
		seq = append(seq, keySeq...)
		p = next
	}
	return seq
}

func typeKeySequence(seq string) string {
	result := make([]uint8, 0)
	p := intVector{0, 0}
	direction := map[uint8]intVector{
		'<': {-1, 0},
		'>': {1, 0},
		'^': {0, -1},
		'v': {0, 1},
	}
	coordinatesToKey := make(map[intVector]uint8)
	for k, v := range keyCoordinates {
		coordinatesToKey[v] = k
	}
	for _, ch := range seq {
		if ch == 'A' {
			result = append(result, coordinatesToKey[p])
		} else {
			p = p.Plus(direction[uint8(ch)])
		}
	}
	return string(result)
}
