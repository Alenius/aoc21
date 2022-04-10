package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type bingoNumber struct {
	number string
	drawn  bool
}

type bingoBoard struct {
	rows  [][]bingoNumber
	bingo bool
}

func (b *bingoBoard) markNumber(number string) {
	for rowIx, row := range b.rows {
		for colIx, val := range row {
			if val.number == number {
				b.rows[rowIx][colIx].drawn = true
				fmt.Printf("number drawn: %v\n", number)
				fmt.Printf("board now: %v\n", b)
			}
		}
	}
}

func (b bingoBoard) checkRowBingo() bool {
	for _, row := range b.rows {
		bingo := true
		for _, val := range row {
			bingo = bingo && val.drawn
		}
		if bingo {
			return bingo
		}
	}

	return false
}

func (b bingoBoard) checkColumnBingo() bool {
	for i := 0; i < 5; i++ {
		bingo := true
		for j := 0; i < 5; i++ {
			val := b.rows[j][i]
			bingo = bingo && val.drawn
			if bingo {
				return bingo
			}
		}
	}

	return false
}

func (b bingoBoard) checkBingo() bool {
	rowBingo := b.checkRowBingo()
	columnBingo := b.checkColumnBingo()
	return rowBingo && columnBingo
}

func readDay4Input(inputPath string) ([]string, []bingoBoard) {
	byteVals, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	vals := string(byteVals)
	splitVals := strings.Split(vals, "\n\n")

	numbersRaw := splitVals[0]
	numbers := strings.Split(numbersRaw, ",")

	boardsRaw := splitVals[1:]

	var boards []bingoBoard
	for _, boardRaw := range boardsRaw {
		board := bingoBoard{rows: [][]bingoNumber{}, bingo: false}
		for rowIx, row := range strings.Split(boardRaw, "\n") {
			board.rows = append(board.rows, []bingoNumber{})
			re := regexp.MustCompile("\\s+")
			for _, val := range re.Split(row, -1) {
				board.rows[rowIx] = append(board.rows[rowIx], bingoNumber{number: val, drawn: false})
			}
		}
		boards = append(boards, board)
	}

	return numbers, boards
}

func (aoc) Day4(inputPath string) {
	numbers, boards := readDay4Input(inputPath)

	fmt.Println(numbers)
	fmt.Println(boards[0])

	winner, winningNumber := drawNumbers(numbers, boards)
	winningSum := findWinningSum(winner)
	winningNumberInt, _ := strconv.Atoi(winningNumber)

	fmt.Printf("day1: %v ", winningSum*winningNumberInt)
}

func drawNumbers(numbers []string, boards []bingoBoard) (bingoBoard, string) {
	for _, number := range numbers {
		for _, board := range boards {
			board.markNumber(number)
			bingo := board.checkBingo()

			if bingo {
				return board, number
			}
		}
	}

	return bingoBoard{}, ""
}

func findWinningSum(board bingoBoard) int {
	sum := 0
	for _, row := range board.rows {
		for _, val := range row {
			if !val.drawn {
				num, _ := strconv.Atoi(val.number)
				sum += num
			}
		}
	}

	return sum
}
