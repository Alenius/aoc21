package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	t "github.com/alenius/aoctools"
)

type signalPatternAndOutput struct {
	signalPatterns []string
	outputs        []string
}

type digitalNumber struct {
	topMid     string
	upperLeft  string
	upperRight string
	midMid     string
	lowerLeft  string
	lowerRight string
	bottomMid  string
}

func (numberMapping digitalNumber) print() {
	fmt.Print("tm: ")
	fmt.Print(numberMapping.topMid)
	fmt.Print(", ul: ")
	fmt.Print(numberMapping.upperLeft)
	fmt.Print(", ur: ")
	fmt.Print(numberMapping.upperRight)
	fmt.Print(", mm: ")
	fmt.Print(numberMapping.midMid)
	fmt.Print(", ll: ")
	fmt.Print(numberMapping.lowerLeft)
	fmt.Print(", lr: ")
	fmt.Print(numberMapping.lowerRight)
	fmt.Print(", bm: ")
	fmt.Println(numberMapping.bottomMid)
}

func (aoc) Day8(inputPath string) {
	lines := t.ReadNewlineSeparatedFile(inputPath)

	patternsAndOutputs := parseInput(lines)

	total := 0
	for _, patternAndOutput := range patternsAndOutputs {
		total += findNoOfOnesFoursSevensAndEights(patternAndOutput.outputs)
	}

	sumOutputs := 0
	for _, patternAndOutput := range patternsAndOutputs {
		numberMappings := findMappingNumbers(patternAndOutput.signalPatterns)
		outputTotal := mapOutputs(numberMappings, patternAndOutput.outputs)
		sumOutputs += outputTotal
	}

	fmt.Print("sum outputs: ")
	fmt.Println(sumOutputs)
}

func parseInput(lines []string) []signalPatternAndOutput {
	var patternsAndOutput []signalPatternAndOutput

	for _, line := range lines {
		splitLine := strings.Split(line, " | ")
		signalPattern := strings.Split(splitLine[0], " ")
		output := strings.Split(splitLine[1], " ")
		patternsAndOutput = append(patternsAndOutput, signalPatternAndOutput{signalPattern, output})
	}

	return patternsAndOutput
}

func findNoOfOnesFoursSevensAndEights(outputs []string) int {
	total := 0

	for _, output := range outputs {
		switch len(output) {
		case 2:
			fallthrough
		case 3:
			fallthrough
		case 4:
			fallthrough
		case 7:
			total++
		}
	}
	return total
}

func findMappingNumbers(signalPatterns []string) digitalNumber {
	var combinationDigitOne []string
	var combinationDigitFour []string
	var combinationDigitSeven []string
	var combinationSixOrNineOrZero [][]string
	var combinationDigitEight []string

	for _, output := range signalPatterns {
		splitOutput := strings.Split(output, "")
		sort.Strings(splitOutput)
		switch len(output) {
		case 2:
			combinationDigitOne = splitOutput
		case 3:
			combinationDigitSeven = splitOutput
		case 4:
			combinationDigitFour = splitOutput
		case 6:
			combinationSixOrNineOrZero = append(combinationSixOrNineOrZero, splitOutput)
		case 7:
			combinationDigitEight = splitOutput
		}
	}

	numberMapping := digitalNumber{}

	sixOrNineOrZero := append(combinationSixOrNineOrZero[0], combinationSixOrNineOrZero[1]...)
	sixOrNineOrZero = append(sixOrNineOrZero, combinationSixOrNineOrZero[2]...)
	urMmLl := removeOverlapping(sixOrNineOrZero, 3)
	ur := findOverlapping(urMmLl, combinationDigitOne)

	numberMapping.upperRight = ur[0]

	urMm := findOverlapping(urMmLl, combinationDigitFour)
	mm := removeOverlapping(append(ur, urMm...), 2)

	numberMapping.midMid = mm[0]
	ll := removeOverlapping(append(urMmLl, urMm...), 2)

	numberMapping.lowerLeft = ll[0]

	lr := removeOverlapping(append(ur, combinationDigitOne...), 2)
	numberMapping.lowerRight = lr[0]

	tm := removeOverlapping(append(combinationDigitSeven, combinationDigitOne...), 2)
	numberMapping.topMid = tm[0]

	oneAndMidMid := append(combinationDigitOne, mm...)
	ul := removeOverlapping(append(oneAndMidMid, combinationDigitFour...), 2)
	numberMapping.upperLeft = ul[0]

	sevenAndMidMid := append(combinationDigitSeven, mm...)
	sevenMmUl := append(sevenAndMidMid, ul...)
	sevenMmUlLl := append(sevenMmUl, ll...)
	bm := removeOverlapping(append(combinationDigitEight, sevenMmUlLl...), 2)
	numberMapping.bottomMid = bm[0]

	return numberMapping
}

func checkForSegment(letters []string, segment string) bool {
	for _, letter := range letters {
		if letter == segment {
			return true
		}
	}

	return false
}

func removeOverlapping(letters []string, numberOfOverlays int) []string {
	letterMap := make(map[string]int)

	for _, letter := range letters {
		letterMap[letter] += 1
	}

	var deduped []string
	for letter, occurences := range letterMap {
		if occurences < numberOfOverlays {
			deduped = append(deduped, letter)
		}
	}

	return deduped
}

func findOverlapping(a []string, b []string) []string {
	allLetters := append(a, b...)

	letterMap := make(map[string]int)

	for _, letter := range allLetters {
		letterMap[letter] += 1
	}

	var deduped []string
	for letter, occurences := range letterMap {
		if occurences > 1 {
			deduped = append(deduped, letter)
		}
	}

	return deduped
}

func mapOutputs(mappings digitalNumber, outputs []string) int {
	var mappedValues [4]int

	for i, output := range outputs {
		splitOutput := strings.Split(output, "")
		switch len(output) {
		case 2:
			mappedValues[i] = 1
		case 3:
			mappedValues[i] = 7
		case 4:
			mappedValues[i] = 4
		case 7:
			mappedValues[i] = 8
		case 6:
			isSix := checkForSegment(splitOutput, mappings.midMid) && checkForSegment(splitOutput, mappings.lowerLeft)
			if isSix {
				mappedValues[i] = 6
				break
			}
			isNine := checkForSegment(splitOutput, mappings.midMid) && checkForSegment(splitOutput, mappings.upperRight)
			if isNine {
				mappedValues[i] = 9
				break
			}

			if !isSix && !isNine {
				mappedValues[i] = 0
			}

		case 5:
			isFive := checkForSegment(splitOutput, mappings.upperLeft) && checkForSegment(splitOutput, mappings.lowerRight)
			if isFive {
				mappedValues[i] = 5
				break
			}
			isTwo := checkForSegment(splitOutput, mappings.lowerLeft) && checkForSegment(splitOutput, mappings.upperRight)
			if isTwo {
				mappedValues[i] = 2
				break
			}
			if !isFive && !isTwo {
				mappedValues[i] = 3
			}
		}
	}

	stringifiedMappedValues := ""
	for _, mappedValue := range mappedValues {
		stringifiedMappedValues += strconv.Itoa(mappedValue)
	}

	mappedOutput, _ := strconv.Atoi(stringifiedMappedValues)

	return mappedOutput
}
