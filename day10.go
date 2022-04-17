package main

import (
	"fmt"
	"sort"

	t "github.com/alenius/aoctools"
)

const (
	oPar = '(' // open parens
	oBrc = '{' // ppen brace
	oBkt = '[' // open bracket
	oGap = '<' // open crocodile gap
	cPar = ')'
	cBrc = '}'
	cBkt = ']'
	cGap = '>'
)

const (
	cParPts    = 3
	cBktPts    = 57
	cBrcPts    = 1197
	cGapPts    = 25137
	cParPtsPt2 = 1
	cBktPtsPt2 = 2
	cBrcPtsPt2 = 3
	cGapPtsPt2 = 4
)

func isOpen(x rune) bool {
	switch x {
	case oPar, oBrc, oBkt, oGap:
		return true
	}
	return false
}

func isClose(x rune) bool {
	switch x {
	case cPar, cBrc, cBkt, cGap:
		return true
	}
	return false
}

func checkCorruptLine(line string) (bool, rune) {
	chunkStack := []rune{}
	for _, char := range line {
		if isOpen(char) {
			chunkStack = append(chunkStack, char)
			continue
		} else if isClose(char) {
			// pop stack
			n := len(chunkStack) - 1
			lastOpen := chunkStack[n]
			chunkStack = chunkStack[:n]

			switch lastOpen {
			case oPar:
				if char != cPar {
					fmt.Printf("expected: %v but found: %v instead\n", string(cPar), string(char))
					return true, char
				}
			case oBrc:
				if char != cBrc {
					fmt.Printf("expected: %v but found: %v instead\n", string(cBrc), string(char))
					return true, char
				}
			case oBkt:
				if char != cBkt {
					fmt.Printf("expected: %v but found: %v instead\n", string(cBkt), string(char))
					return true, char
				}
			case oGap:
				if char != cGap {
					fmt.Printf("expected: %v but found: %v instead\n", string(cGap), string(char))
					return true, char
				}

			default:
				fmt.Println("no of the above")
			}

		} else {
			fmt.Println("not recognized char")
		}

	}
	return false, 0
}

func countCorruptPoints(ccs []rune) int {
	var pts int
	for _, cc := range ccs {
		fmt.Println(string(cc))
		switch cc {
		case cPar:
			pts += cParPts
		case cBkt:
			pts += cBktPts
		case cBrc:
			pts += cBrcPts
		case cGap:
			pts += cGapPts
		default:
			fmt.Println("no idea what char that is")
		}
	}

	return pts
}

func reverseSlice(slice []rune) []rune {
	reversedSlice := []rune{}
	for ix := range slice {
		reversedSlice = append(reversedSlice, slice[len(slice)-1-ix])
	}
	return reversedSlice
}

func countIncompletePoints(lines [][]rune) int {
	var pts []int

	for _, line := range lines {
		innerPts := 0
		// stack gives us wrong order so need to reverse it
		reversed := reverseSlice(line)
		for _, cc := range reversed {
			innerPts = innerPts * 5
			switch cc {
			case cPar:
				innerPts += cParPtsPt2
			case cBkt:
				innerPts += cBktPtsPt2
			case cBrc:
				innerPts += cBrcPtsPt2
			case cGap:
				innerPts += cGapPtsPt2
			default:
				fmt.Println("no idea what char that is")
			}
		}
		fmt.Printf("pts: %v \n", innerPts)
		pts = append(pts, innerPts)
	}

	sort.Ints(pts)

	val := pts[(len(pts)-1)/2]

	return val
}

func (aoc) Day10(inputPath string) {
	lines := t.ReadNewlineSeparatedFile(inputPath)

	corruptLines := []string{}
	incompleteLines := []string{}
	corruptFirstChar := []rune{}
	for _, line := range lines {
		isCorrupt, firstCorruptChar := checkCorruptLine(line)
		if isCorrupt {
			corruptLines = append(corruptLines, line)
			corruptFirstChar = append(corruptFirstChar, firstCorruptChar)
		} else {
			incompleteLines = append(incompleteLines, line)
		}

	}

	fmt.Printf("corrupt lines: %+v\n", corruptLines[0])
	fmt.Printf("corrupt lines len: %+v\n", len(corruptLines))

	pts := countCorruptPoints(corruptFirstChar)
	fmt.Printf("corrupt pts: %v\n", pts)
	completionChars := completeIncompleteLines(incompleteLines)
	ptsPt2 := countIncompletePoints(completionChars)
	fmt.Printf("incomplete pts: %v\n", ptsPt2)
}

func completeIncompleteLines(lines []string) [][]rune {
	remainingOpenChunks := [][]rune{}
	for _, line := range lines {
		chunkStack := []rune{}
		for _, char := range line {
			if isOpen(char) {
				chunkStack = append(chunkStack, char)
				continue
			} else if isClose(char) {
				// pop stack
				n := len(chunkStack) - 1
				chunkStack = chunkStack[:n]
			} else {
				fmt.Println("not recognized char")
			}
		}

		remainingOpenChunks = append(remainingOpenChunks, chunkStack)
	}

	completionChars := [][]rune{}
	// all done, check remaining open
	for _, line := range remainingOpenChunks {
		innerCompletionChars := []rune{}
		for _, remaining := range line {
			switch remaining {
			case oPar:
				innerCompletionChars = append(innerCompletionChars, cPar)
			case oBkt:
				innerCompletionChars = append(innerCompletionChars, cBkt)
			case oBrc:
				innerCompletionChars = append(innerCompletionChars, cBrc)
			case oGap:
				innerCompletionChars = append(innerCompletionChars, cGap)
			default:
				fmt.Println("do not regocize")
			}
		}
		completionChars = append(completionChars, innerCompletionChars)
	}

	for _, il := range completionChars[4] {
		fmt.Printf("il %v\n", string(il))
	}

	return completionChars
}
