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
	//proposeWrongCandidateSets(orderedGates, memoryOffsets)
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

func proposeWrongCandidateSets(orderedGates []gateData, memoryOffsets map[string]int) {
	outputToGate := make(map[string]gateData)
	for _, gate := range orderedGates {
		outputToGate[gate.output] = gate
	}
	outputs := make([]int, len(memoryOffsets))
	for i := 0; i < len(memoryOffsets); i++ {
		outputs[i] = i
	}

	wrongSetCandidates := make(map[string]bool)
	/*	for maxBit := 14; maxBit < 15; maxBit++ {
			fmt.Printf("maxBit=%d\n", maxBit)
			for x := 1 << maxBit; x < 1<<(maxBit+1); x++ {
				//fmt.Printf("x=%d\n", x)
				for y := 0; y <= x; y++ {
					if compiledProgram(x, y, outputs) != x+y {
						changedNames := make([]string, 0)
						for b := 0; b <= maxBit; b++ {
							if 1<<b&x != 0 {
								changedNames = append(changedNames, fmt.Sprintf("x%02d", b))
							}
							if 1<<b&y != 0 {
								changedNames = append(changedNames, fmt.Sprintf("y%02d", b))
							}
							v := createInitialValues(changedNames)
							if !checkAndCollectWrongSetCandidates(orderedGates, outputToGate, v, changedNames, wrongSetCandidates) {
								//fmt.Printf("{%d, %d},\n", xi, yi)
							}
						}
					}
				}
			}
		}
	*/
	for yi := 0; yi < totalNumOfBits; yi++ {
		for xi := 0; xi < totalNumOfBits; xi++ {
			changedNames := []string{
				fmt.Sprintf("x%02d", xi),
				fmt.Sprintf("y%02d", yi),
			}
			v := createInitialValues(changedNames)
			if !checkAndCollectWrongSetCandidates(orderedGates, outputToGate, v, changedNames, wrongSetCandidates) {
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
	allCandidates := Map(gates, func(gate gateData) string { return gate.output })
	initialCandidates := []string{
		"ncd", "nfj",
		"z15", "qnw",
		"z20", "cqr",
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
}

func findCorrectionsForCandidates(memoryOffsets map[string]int, candidates []string) {
	corrections := candidates
	errorIndex := checkCorrectionsAndReturnInvalidIndex(memoryOffsets, corrections)
	if errorIndex == -1 {
		fmt.Println("Found!!!!")
		fmt.Printf("Corrections: %s<->%s, %s<->%s, %s<->%s, %s<->%s\n",
			corrections[0], corrections[1], corrections[2], corrections[3], corrections[4], corrections[5],
			corrections[6], corrections[7],
		)
		sorted := slices.Clone(corrections)
		slices.Sort(sorted)
		fmt.Printf("Answer: %s", strings.Join(sorted, ","))
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

func checkAndCollectWrongSetCandidates(orderedGates []gateData, outputToGate map[string]gateData, values map[string]int, changedNames []string, newWrongCandidatesSet map[string]bool) bool {
	corrections := make(map[string]string)
	corrections["z20"] = "cqr"
	corrections["cqr"] = "z20"
	corrections["qnw"] = "z15"
	corrections["z15"] = "qnw"
	corrections["ncd"] = "nfj"
	corrections["nfj"] = "ncd"
	runComputation(orderedGates, values, corrections)
	x := bitsToValue("x", values)
	y := bitsToValue("y", values)
	z := bitsToValue("z", values)
	if x+y == z {
		return true
	}
	changedNamesSet := make(map[string]bool)
	for _, name := range changedNames {
		changedNamesSet[name] = true
	}
	expectedBits := valueToBits("z", x+y, totalNumOfBits+1)
	for name, value := range expectedBits {
		if value != values[name] {
			mayCause, brokenCandidates := mayCauseIncorrectValue(name, values[name], values, outputToGate)
			if mayCause {
				candidatesSlice := Keys(brokenCandidates)
				slices.Sort(candidatesSlice)
				candidatesString := strings.Join(candidatesSlice, ",")
				if !newWrongCandidatesSet[candidatesString] {
					newWrongCandidatesSet[candidatesString] = true
					fmt.Printf("Incorrect computation for %s: %s\n", strings.Join(changedNames, ","), candidatesString)
				}
			}
		}
	}
	return false
}

func mayCauseIncorrectValue(name string, incorrectValue int, values map[string]int, outputToGate map[string]gateData) (bool, map[string]bool) {
	if strings.HasPrefix(name, "x") || strings.HasPrefix(name, "y") {
		return values[name] == incorrectValue, make(map[string]bool)
	}
	gate, ok := outputToGate[name]
	if !ok {
		panic("Cannot find gate for output " + name)
	}
	switch gate.operationSign {
	case "&":
		if incorrectValue == 1 {
			i1, s1 := mayCauseIncorrectValue(gate.input1, 1, values, outputToGate)
			i2, s2 := mayCauseIncorrectValue(gate.input2, 1, values, outputToGate)
			if i1 && i2 {
				return true, AddElement(Union(s1, s2), gate.output)
			} else {
				return false, nil
			}
		} else {
			i1, s1 := mayCauseIncorrectValue(gate.input1, 0, values, outputToGate)
			i2, s2 := mayCauseIncorrectValue(gate.input2, 0, values, outputToGate)
			if i1 && i2 {
				return true, AddElement(Union(s1, s2), gate.output)
			} else if i1 {
				return true, AddElement(s1, gate.output)
			} else if i2 {
				return true, AddElement(s2, gate.output)
			} else {
				return false, nil
			}
		}
	case "|":
		if incorrectValue == 0 {
			i1, s1 := mayCauseIncorrectValue(gate.input1, 0, values, outputToGate)
			i2, s2 := mayCauseIncorrectValue(gate.input2, 0, values, outputToGate)
			if i1 && i2 {
				return true, AddElement(Union(s1, s2), gate.output)
			} else {
				return false, nil
			}
		} else {
			i1, s1 := mayCauseIncorrectValue(gate.input1, 1, values, outputToGate)
			i2, s2 := mayCauseIncorrectValue(gate.input2, 1, values, outputToGate)
			if i1 && i2 {
				return true, AddElement(Union(s1, s2), gate.output)
			} else if i1 {
				return true, AddElement(s1, gate.output)
			} else if i2 {
				return true, AddElement(s2, gate.output)
			} else {
				return false, nil
			}
		}
	case "^":
		i10, s10 := mayCauseIncorrectValue(gate.input1, 0, values, outputToGate)
		i20, s20 := mayCauseIncorrectValue(gate.input2, 0, values, outputToGate)
		i11, s11 := mayCauseIncorrectValue(gate.input1, 1, values, outputToGate)
		i21, s21 := mayCauseIncorrectValue(gate.input2, 1, values, outputToGate)
		if incorrectValue == 1 {
			if i10 && i21 && i11 && i20 {
				return true, AddElement(Union(Union(s10, s21), Union(s11, s20)), gate.output)
			} else if i10 && i21 {
				return true, AddElement(Union(s10, s21), gate.output)
			} else if i11 && i20 {
				return true, AddElement(Union(s11, s20), gate.output)
			} else {
				return false, nil
			}
		} else {
			if i10 && i20 && i11 && i21 {
				return true, AddElement(Union(Union(s10, s20), Union(s11, s21)), gate.output)
			} else if i10 && i20 {
				return true, AddElement(Union(s10, s20), gate.output)
			} else if i11 && i21 {
				return true, AddElement(Union(s11, s21), gate.output)
			} else {
				return false, nil
			}
		}
	default:
		panic("Unknown operation sign " + gate.operationSign)
	}
	return false, nil
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

// , , z20<->cqr,
var wrongCandidateSets = [][]string{
	{"dnt", "z37"}, //vkg<->z37
	{"ctg", "gnc", "kdf", "qnw", "rjm", "z16"},        //z15<->qnw
	{"cqr", "hbp", "msn", "rjq", "vct", "wrb", "z21"}, //ncd<->nfj
	{"brm", "kmn", "mdv", "ncd", "ndd", "qrs", "z29"},
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
