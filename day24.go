package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func day24(fileName string) {
	lines := ReadLines(fileName)
	values := make(map[string]int)
	assignmentRegexp := regexp.MustCompile("(\\w+): ([0-1])")
	gateRegexp := regexp.MustCompile("(\\w+) (\\w+) (\\w+) -> (\\w+)")
	gates := make([]gateData, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		assignment := assignmentRegexp.FindStringSubmatch(line)
		if assignment != nil {
			values[assignment[1]] = StrToInt(assignment[2])
		} else {
			gate := gateRegexp.FindStringSubmatch(line)
			if gate == nil {
				panic("unexpected format " + line)
			}
			var operation func(int, int) int
			switch gate[2] {
			case "AND":
				operation = andFunction
			case "OR":
				operation = orFunction
			case "XOR":
				operation = xorFunction
			default:
				panic("unknown gate type " + gate[2])
			}
			gates = append(gates, gateData{
				gate[1],
				gate[3],
				gate[4],
				operation,
			})
		}
	}
	changed := true
	for changed {
		changed = false
		for _, gate := range gates {
			input1, ok1 := values[gate.input1]
			input2, ok2 := values[gate.input2]
			_, ok3 := values[gate.output]
			if ok1 && ok2 && !ok3 {
				output := gate.operation(input1, input2)
				values[gate.output] = output
				changed = true
			}
		}
	}
	zKeys := make([]string, 0)
	for key, _ := range values {
		if strings.HasPrefix(key, "z") {
			zKeys = append(zKeys, key)
		}
	}
	sort.Strings(zKeys)
	result := 0
	for i := len(zKeys) - 1; i >= 0; i-- {
		result = 2*result + values[zKeys[i]]
	}
	fmt.Println(result)
}

type gateData struct {
	input1    string
	input2    string
	output    string
	operation func(int, int) int
}

var andFunction = func(a, b int) int { return a & b }
var orFunction = func(a, b int) int { return a | b }
var xorFunction = func(a, b int) int { return a ^ b }
