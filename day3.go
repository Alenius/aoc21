package main

import (
	"log"
	"strconv"

	t "github.com/alenius/aoctools"
)

type GammaEpsilon struct {
	gamma   int
	epsilon int
}

func readDay3Input(lines []string) []GammaEpsilon {
	var gammaEpsilons []GammaEpsilon
	bitwidth := len(lines[0])

	for len(gammaEpsilons) < bitwidth {
		noOfOnes := 0
		noOfZeros := 0

		for _, line := range lines {
			value := line[len(gammaEpsilons)]
			if value == '1' {
				noOfOnes++
			} else {
				noOfZeros++
			}
		}

		var gamma int
		var epsilon int
		if noOfOnes > noOfZeros {
			gamma = 1
			epsilon = 0
		} else {
			gamma = 0
			epsilon = 1
		}

		gammaEpsilons = append(gammaEpsilons, GammaEpsilon{gamma, epsilon})
	}

	return gammaEpsilons
}

type O2CO2Rating struct {
	ox   int
	cdox int
}

func solveDay2(lines []string, keepMost bool) int {
	ix := 0
	remaningLines := lines

	for {

		if len(remaningLines) <= 1 || ix >= len(lines[0]) {
			break
		}

		noOfOnes := 0
		noOfZeros := 0

		for _, line := range remaningLines {
			value := line[ix]
			if value == '1' {
				noOfOnes++
			} else {
				noOfZeros++
			}
		}

		mostOnes := noOfOnes >= noOfZeros
		var linesToKeep []string
		for _, line := range remaningLines {
			if keepMost {
				if mostOnes {
					if line[ix] == '1' {
						linesToKeep = append(linesToKeep, line)
					}
				} else {
					if line[ix] == '0' {
						linesToKeep = append(linesToKeep, line)
					}
				}
			} else {
				if mostOnes {
					if line[ix] == '0' {
						linesToKeep = append(linesToKeep, line)
					}
				} else {
					if line[ix] == '1' {
						linesToKeep = append(linesToKeep, line)
					}
				}

			}
		}

		remaningLines = linesToKeep
		ix++
	}

	final_value, _ := strconv.ParseInt(remaningLines[0], 2, 64)
	return int(final_value)
}

func (aoc) Day3(inputPath string) {
	lines := t.ReadNewlineSeparatedFile(inputPath)
	gammaEpsilons := readDay3Input(lines)

	log.Println("consumption", gammaEpsilons)

	var gammaString string
	var epsilonString string
	for _, gammaEpsilon := range gammaEpsilons {
		gammaStr := strconv.Itoa(gammaEpsilon.gamma)
		epsilonStr := strconv.Itoa(gammaEpsilon.epsilon)
		gammaString = gammaString + gammaStr
		epsilonString = epsilonString + epsilonStr
	}

	gammaRate, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilonRate, _ := strconv.ParseInt(epsilonString, 2, 64)

	log.Println("pt1", gammaRate*epsilonRate)

	// pt2
	ox := solveDay2(lines, true)
	co2 := solveDay2(lines, false)

	log.Println(ox)
	log.Println(co2)

	log.Println("pt2", ox*co2)
}
