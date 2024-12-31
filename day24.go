package main

import (
	"fmt"
	"regexp"
	"slices"
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
			var operationSign string
			switch gate[2] {
			case "AND":
				operation = andFunction
				operationSign = "&"
			case "OR":
				operation = orFunction
				operationSign = "|"
			case "XOR":
				operation = xorFunction
				operationSign = "^"
			default:
				panic("unknown gate type " + gate[2])
			}
			gates = append(gates, gateData{
				gate[1],
				gate[3],
				gate[4],
				operationSign,
				operation,
			})
		}
	}
	outputs := make([]int, 312)
	for i := 0; i < len(outputs); i++ {
		outputs[i] = i
	}
	memoryOffsets, orderedGates := computeMemoryOffsetsAndOrderedGates(gates)
	_ = orderedGates
	_ = memoryOffsets
	//compileProgram(orderedGates, memoryOffsets)
	//proposeWrongCandidateSets(gates)
	findCorrections(memoryOffsets, orderedGates)
}

func computeMemoryOffsetsAndOrderedGates(gates []gateData) (map[string]int, []gateData) {
	offsets := make(map[string]int)
	nextOffset := 0
	computed := make(map[string]bool)
	for x := 0; x < totalNumOfBits; x++ {
		name := fmt.Sprintf("x%02d", x)
		offsets[name] = nextOffset
		computed[name] = true
		nextOffset++
	}
	for y := 0; y < totalNumOfBits; y++ {
		name := fmt.Sprintf("y%02d", y)
		offsets[name] = nextOffset
		computed[name] = true
		nextOffset++
	}
	changed := true
	orderedGates := make([]gateData, 0)
	for changed {
		changed = false
		for _, gate := range gates {
			if computed[gate.input1] && computed[gate.input2] && !computed[gate.output] {
				offsets[gate.output] = nextOffset
				computed[gate.output] = true
				orderedGates = append(orderedGates, gate)
				changed = true
				nextOffset++
			}
		}
	}
	return offsets, orderedGates
}

func compileProgram(gates []gateData, offsets map[string]int) {
	fmt.Println("func compiledProgram(x int, y int, outputs []int) int {")
	fmt.Printf("  memory := [%d]int{}\n", len(offsets))
	fmt.Println("  changed := true")
	fmt.Printf("  for i := 0; i < %d; i++ {\n", len(offsets))
	fmt.Println("    memory[i] = -1")
	fmt.Println("  }")
	for i := 0; i < totalNumOfBits; i++ {
		fmt.Printf("  memory[%d] = (x >> %d) & 1\n", offsets[fmt.Sprintf("x%02d", i)], i)
		fmt.Printf("  memory[%d] = (y >> %d) & 1\n", offsets[fmt.Sprintf("y%02d", i)], i)
	}
	fmt.Println("  for i := 0; i < 8 && changed; i++ {")
	fmt.Println("    changed = false")
	for _, gate := range gates {
		input1 := offsets[gate.input1]
		input2 := offsets[gate.input2]
		output := offsets[gate.output]
		fmt.Printf("    if memory[%d] != -1 && memory[%d] != -1 && memory[outputs[%d]] == -1 {\n", input1, input2, output)
		fmt.Println("      changed = true")
		fmt.Printf("      memory[outputs[%d]] = memory[%d] %s memory[%d]\n", output, input1, gate.operationSign, input2)
		fmt.Println("    }")
	}
	fmt.Println("  }")
	fmt.Println("  if changed {")
	fmt.Println("    panic(\"failed to assign all values\")")
	fmt.Println("  }")
	fmt.Println("  z := 0")
	for i := 0; i < totalNumOfBits+1; i++ {
		fmt.Printf("  z |= memory[%d] << %d\n", offsets[fmt.Sprintf("z%02d", i)], i)
	}
	fmt.Println("  return z")
	fmt.Println("}")
}

func proposeWrongCandidateSets(gates []gateData) {
	wrongSetCandidates := make(map[string]bool)
	for yi := 0; yi < totalNumOfBits; yi++ {
		for xi := 0; xi < totalNumOfBits; xi++ {
			changedNames := []string{
				fmt.Sprintf("x%02d", xi),
				fmt.Sprintf("y%02d", yi),
			}
			v := createInitialValues(changedNames)
			if !checkAndCollectWrongSetCandidates(gates, v, changedNames, wrongSetCandidates) {
				fmt.Printf("{%d, %d},\n", xi, yi)
			}
		}
	}
}

func createInitialValues(changedNames []string) map[string]int {
	v := make(map[string]int)
	for i := 0; i < totalNumOfBits; i++ {
		v[fmt.Sprintf("x%02d", i)] = 0
		v[fmt.Sprintf("y%02d", i)] = 0
	}
	for _, name := range changedNames {
		v[name] = 1
	}
	return v
}

func findCorrections(memoryOffsets map[string]int, gates []gateData) {
	candidateInSet := make([]int, len(wrongCandidateSets))
	allCandidates := Map(gates, func(gate gateData) string { return gate.output })
	for {
		initialCandidates := make([]string, len(wrongCandidateSets))
		for i, index := range candidateInSet {
			initialCandidates[i] = wrongCandidateSets[i][index]
		}
		for i5 := 0; i5 < len(allCandidates); i5++ {
			c5 := allCandidates[i5]
			if slices.Contains(initialCandidates, c5) {
				continue
			}
			for i6 := i5 + 1; i6 < len(allCandidates); i6++ {
				c6 := allCandidates[i6]
				if slices.Contains(initialCandidates, c6) {
					continue
				}
				candidates := slices.Concat(initialCandidates, []string{c5, c6})
				findCorrectionsForCandidates(memoryOffsets, candidates)
			}
		}
		index := len(candidateInSet) - 1
		for index >= 0 {
			if candidateInSet[index] < len(wrongCandidateSets[index])-1 {
				candidateInSet[index]++
				break
			} else {
				index--
			}
		}
		if index < 0 {
			break
		}
	}
}

func findCorrectionsForCandidates(memoryOffsets map[string]int, candidates []string) {
	bestPermutation := ""
	maxErrorIndex := 0
	c0a := 0
	for c0b := 1; c0b <= 5; c0b++ {
		var c1a int
		if c0b == 1 {
			c1a = 2
		} else {
			c1a = 1
		}
		for c1b := 1; c1b <= 5; c1b++ {
			if c1b == c0b || c1b == c1a {
				continue
			}
			c2a := c1a + 1
			for c2a == c0b || c2a == c1b {
				c2a++
			}
			c2b := c2a + 1
			for c2b == c0b || c2b == c1b {
				c2b++
			}
			corrections := []string{
				candidates[c0a], candidates[c0b],
				candidates[c1a], candidates[c1b],
				candidates[c2a], candidates[c2b],
				"z20", "cqr",
			}
			errorIndex := checkCorrectionsAndReturnInvalidIndex(memoryOffsets, corrections)
			if errorIndex == -1 {
				fmt.Println("Found!!!!")
				fmt.Printf("Corrections: %s<->%s, %s<->%s, %s<->%s, z20<->cqr\n",
					candidates[c0a], candidates[c0b], candidates[c1a], candidates[c1b],
					candidates[c2a], candidates[c2b])
			}
			if errorIndex > maxErrorIndex {
				bestPermutation = fmt.Sprintf("Corrections: %s<->%s, %s<->%s, %s<->%s, z20<->cqr",
					candidates[c0a], candidates[c0b], candidates[c1a], candidates[c1b],
					candidates[c2a], candidates[c2b])
				maxErrorIndex = errorIndex
			}
		}
	}
	if maxErrorIndex >= 27 {
		//fmt.Printf("Checking corrections for candidates: %v...", candidates)
		fmt.Printf("For %s max error index = %d\n", bestPermutation, maxErrorIndex)
	}
}

func checkCorrectionsAndReturnInvalidIndex(memoryOffsets map[string]int, corrections []string) int {
	errorIndex := 0
	outputs := make([]int, len(memoryOffsets))
	for i := 0; i < len(memoryOffsets); i++ {
		outputs[i] = i
	}
	for i := 0; i < len(corrections); i += 2 {
		o1 := memoryOffsets[corrections[i]]
		o2 := memoryOffsets[corrections[i+1]]
		outputs[o1] = o2
		outputs[o2] = o1
	}
	for _, xy := range xyIndexesToCheckFirst {
		if !checkCorrectionsForChangedValues(xy[0], xy[1], outputs) {
			return errorIndex
		}
		errorIndex++
	}
	for xi := 0; xi < totalNumOfBits; xi++ {
		for yi := 0; yi < totalNumOfBits; yi++ {
			if !checkCorrectionsForChangedValues(xi, yi, outputs) {
				return errorIndex
			}
			errorIndex++
		}
	}
	return -1
}

func checkCorrectionsForChangedValues(xi int, yi int, outputs []int) bool {
	x := 1 << xi
	y := 1 << yi
	z := compiledProgram(x, y, outputs)
	return x+y == z
}

func checkAndCollectWrongSetCandidates(gates []gateData, values map[string]int, changedNames []string, newWrongCandidatesSet map[string]bool) bool {
	corrections := make(map[string]string)
	corrections["z20"] = "cqr"
	corrections["cqr"] = "z20"
	runComputation(gates, values, corrections)
	x := bitsToValue("x", values)
	y := bitsToValue("y", values)
	z := bitsToValue("z", values)
	if x+y == z {
		return true
	}
	expectedBits := valueToBits("z", x+y, totalNumOfBits+1)
	incorrectBits := make(map[string]bool)
	for name, value := range expectedBits {
		if value != values[name] {
			incorrectBits[name] = true
		}
	}
	brokenCandidates := make(map[string]bool)
	for _, changedName := range changedNames {
		computeAffectedNodes(changedName, incorrectBits, gates, brokenCandidates)
	}
	if Any(wrongCandidateSets, func(set []string) bool {
		return Any(set, func(s string) bool {
			return brokenCandidates[s]
		})
	}) {
		return false
	}
	candidatesSlice := Keys(brokenCandidates)
	slices.Sort(candidatesSlice)
	candidatesString := strings.Join(candidatesSlice, ",")
	if !newWrongCandidatesSet[candidatesString] {
		newWrongCandidatesSet[candidatesString] = true
		//fmt.Printf("Incorrect computation for %s: %s\n", strings.Join(changedNames, ","), candidatesString)
	}
	return false
}

func computeAffectedNodes(name string, incorrectBits map[string]bool, gates []gateData, candidates map[string]bool) bool {
	if incorrectBits[name] {
		return true
	}
	for _, gate := range gates {
		if gate.input1 == name || gate.input2 == name {
			if computeAffectedNodes(gate.output, incorrectBits, gates, candidates) {
				candidates[gate.output] = true
				return true
			}
		}
	}
	return false
}

func runComputation(gates []gateData, values map[string]int, outputCorrections map[string]string) {
	changed := true
	for changed {
		changed = false
		for _, gate := range gates {
			input1, ok1 := values[gate.input1]
			input2, ok2 := values[gate.input2]
			actualOutput, ok := outputCorrections[gate.output]
			if !ok {
				actualOutput = gate.output
			}
			_, ok3 := values[actualOutput]
			if ok1 && ok2 && !ok3 {
				output := gate.operation(input1, input2)
				values[actualOutput] = output
				changed = true
			}
		}
	}
}

func bitsToValue(name string, values map[string]int) int {
	result := 0
	for i := totalNumOfBits - 1; i >= 0; i-- {
		result = 2*result + values[fmt.Sprintf("%s%02d", name, i)]
	}
	return result
}

func valueToBits(name string, value int, bits int) map[string]int {
	result := make(map[string]int)
	for i := 0; i < bits; i++ {
		result[fmt.Sprintf("%s%02d", name, i)] = (value >> i) & 1
	}
	return result
}

var wrongCandidateSets = [][]string{
	{"ncd", "ndd", "z28"},
	{"qnw", "rjm", "z16"},
	{"bvv", "dnt", "vkg", "z38"},
	{"fkg", "hbp", "vct", "wrb", "z21"},
}

var totalNumOfBits = 45

type gateData struct {
	input1        string
	input2        string
	output        string
	operationSign string
	operation     func(int, int) int
}

var andFunction = func(a, b int) int { return a & b }
var orFunction = func(a, b int) int { return a | b }
var xorFunction = func(a, b int) int { return a ^ b }
