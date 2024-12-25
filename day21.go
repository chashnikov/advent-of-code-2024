package main

import (
	"fmt"
	"slices"
	"strings"
)

func day21(fileName string) {
	lines := ReadLines(fileName)
	cache := make([]map[string]int, 30)
	for device := 0; device < len(cache); device++ {
		cache[device] = make(map[string]int)
	}
	fmt.Println(Sum(Map(lines, func(line string) int {
		num := StrToInt(strings.TrimSuffix(line, "A"))
		keySequenceLen := ComputeKeySequenceLen([]uint8(line), 0, cache)
		fmt.Print(line)
		fmt.Print(":")
		fmt.Println(keySequenceLen)
		return num * keySequenceLen
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

func ComputeKeySequenceLen(line []uint8, device int, cache []map[string]int) int {
	if device == 26 {
		return len(line)
	}
	p := intVector{0, 0}
	cached, ok := cache[device][string(line)]
	if ok {
		return cached
	}
	seqLen := 0
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
		keySeqLen := 0
		if p.Plus(intVector{dx, 0}) != forbiddenCoordinate {
			keySeqLen = ComputeKeySequenceLen(slices.Concat(horizontal, vertical, []uint8{'A'}), device+1, cache)
		}
		if p.Plus(intVector{0, dy}) != forbiddenCoordinate {
			newKeySeq := ComputeKeySequenceLen(slices.Concat(vertical, horizontal, []uint8{'A'}), device+1, cache)
			if keySeqLen == 0 || keySeqLen > newKeySeq {
				keySeqLen = newKeySeq
			}
		}
		seqLen += keySeqLen
		p = next
	}
	if p.x != 0 || p.y != 0 {
		panic("key sequence doesn't finish at the origin")
	}
	cache[device][string(line)] = seqLen
	return seqLen
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
