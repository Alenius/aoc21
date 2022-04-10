package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ventCoordinates struct {
	fromX int
	fromY int
	toX   int
	toY   int
}

func parseDay5Input(inputPath string) []ventCoordinates {
	file, err := os.Open(inputPath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var coordinates []ventCoordinates
	for scanner.Scan() {
		newLine := scanner.Text()
		coordTuple := strings.Split(newLine, "->")

		from := strings.TrimSpace(coordTuple[0])
		to := strings.TrimSpace(coordTuple[1])

		fromSplit := strings.Split(from, ",")
		toSplit := strings.Split(to, ",")

		fromX, _ := strconv.Atoi(fromSplit[0])
		fromY, _ := strconv.Atoi(fromSplit[1])
		toX, _ := strconv.Atoi(toSplit[0])
		toY, _ := strconv.Atoi(toSplit[1])

		coordinate := ventCoordinates{fromX, fromY, toX, toY}
		coordinates = append(coordinates, coordinate)
	}

	return coordinates
}

func (aoc) Day5(inputPath string) {
	coordinates := parseDay5Input(inputPath)

	ventPoints := make(map[[2]int]int)

	// pt1
	overlaps := solveStraightHorizontalOrVertical(coordinates, ventPoints)
	log.Println("PT1", overlaps)

	// pt2
	overlapsCountingDiagonal := solveDiagonalLines(coordinates, ventPoints)
	log.Println("PT2", overlapsCountingDiagonal)
}

func solveDiagonalLines(coordinates []ventCoordinates, ventPoints map[[2]int]int) int {
	for _, coordinate := range coordinates {
		isDiagonal := checkDiagonality(coordinate.fromX, coordinate.fromY, coordinate.toX, coordinate.toY)
		log.Println("is diagonal", isDiagonal)
		log.Println("point", coordinate)

		if !isDiagonal {
			continue
		}

		newPoints := findDiagonalDiffs(coordinate)

		if len(newPoints) > 0 {
			for _, newPoint := range newPoints {
				findOverlaps(ventPoints, newPoint[0], newPoint[1])
			}
		}

	}

	numberOfOverlaps := 0
	log.Println("vent points", ventPoints)
	for _, elem := range ventPoints {
		if elem > 1 {
			numberOfOverlaps++
		}
	}

	log.Println("nubmer of overlaps", numberOfOverlaps)

	return numberOfOverlaps
}

func solveStraightHorizontalOrVertical(coordinates []ventCoordinates, ventPoints map[[2]int]int) int {

	for _, coordinate := range coordinates {
		horizontalStraight := coordinate.fromX == coordinate.toX
		verticalStraight := coordinate.fromY == coordinate.toY

		if !horizontalStraight && !verticalStraight {
			continue
		}

		diffX := absDiffInt(coordinate.fromX, coordinate.toX)

		if diffX != 0 {
			newPoints := findDiffs(true, diffX, coordinate)

			for _, newPoint := range newPoints {
				findOverlaps(ventPoints, newPoint[0], newPoint[1])
			}
		}

		diffY := absDiffInt(coordinate.fromY, coordinate.toY)
		if diffY != 0 {
			newPoints := findDiffs(false, diffY, coordinate)

			for _, newPoint := range newPoints {
				findOverlaps(ventPoints, newPoint[0], newPoint[1])
			}
		}
	}

	numberOfOverlaps := 0
	for _, elem := range ventPoints {
		if elem > 1 {
			numberOfOverlaps++
		}
	}

	log.Println("nubmer of overlaps", numberOfOverlaps)
	return numberOfOverlaps
}

func checkDiagonality(fromX, fromY, toX, toY int) bool {
	diffX := absDiffInt(fromX, toX)
	diffY := absDiffInt(fromY, toY)
	return diffX == diffY
	// diagonalX := fromX == toY
	// diagonalY := fromY == toX

	// symmetrical := (fromX == fromY) && (toX == toY)

	// return (diagonalX && diagonalY) || symmetrical
}

func findDiffs(xVals bool, diffValue int, coordinate ventCoordinates) [][2]int {
	var diffs []([2]int)

	if xVals {
		for i := 0; i <= diffValue; i++ {
			var newX int
			if coordinate.fromX < coordinate.toX {
				newX = coordinate.fromX + i
			} else {
				newX = coordinate.fromX - i
			}

			x := newX
			y := coordinate.fromY

			newPoint := [2]int{x, y}
			diffs = append(diffs, newPoint)
		}
	} else {
		for i := 0; i <= diffValue; i++ {
			var newY int
			if coordinate.fromY < coordinate.toY {
				newY = coordinate.fromY + i
			} else {
				newY = coordinate.fromY - i
			}

			x := coordinate.fromX
			y := newY

			newPoint := [2]int{x, y}
			diffs = append(diffs, newPoint)

		}
	}
	return diffs
}

func findDiagonalDiffs(coordinate ventCoordinates) [][2]int {
	var diffs []([2]int)

	fromX := coordinate.fromX
	fromY := coordinate.fromY
	toX := coordinate.toX
	toY := coordinate.toY

	positiveDiffX := (toX - fromX) > 0
	positiveDiffY := (toY - fromY) > 0

	distance := absDiffInt(fromX, toX)

	for i := 0; i <= distance; i++ {
		var newX int
		var newY int
		if positiveDiffX && positiveDiffY {
			newX = fromX + i
			newY = fromY + i
		} else if positiveDiffX && !positiveDiffY {
			newX = fromX + i
			newY = fromY - i
		} else if !positiveDiffX && positiveDiffY {
			newX = fromX - i
			newY = fromY + i
		} else {
			newX = fromX - i
			newY = fromY - i
		}

		newPoint := [2]int{newX, newY}
		log.Println("new point", newPoint)
		diffs = append(diffs, newPoint)
	}

	return diffs
}

func findOverlaps(ventPoints map[[2]int]int, x, y int) {
	mapKey := [2]int{x, y}

	currentOverlaps := ventPoints[mapKey]

	if currentOverlaps != 0 {
		ventPoints[mapKey] = currentOverlaps + 1
	} else {
		ventPoints[mapKey] = 1
	}

}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}
