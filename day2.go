package main

import (
	"log"
	"strconv"
	"strings"

	t "github.com/alenius/aoctools"
)

type PositionChange struct {
	value     int
	direction string
}

type Position struct {
	x   int
	y   int
	aim int
}

func readDay2Input(lines []string) []PositionChange {
	var positionChanges []PositionChange
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		direction := splitLine[0]
		rawValue := splitLine[1]
		value, _ := strconv.Atoi(rawValue)
		positionChanges = append(positionChanges, PositionChange{value, direction})
	}

	return positionChanges
}

func (aoc) Day2(inputPath string) {
	lines := t.ReadNewlineSeparatedFile(inputPath)
	measurements := readDay2Input(lines)

	log.Println(len(lines))

	// pt 1
	var position Position

	for _, measurement := range measurements {
		switch measurement.direction {
		case "forward":
			position.x = position.x + measurement.value
		case "backward":
			position.x = position.x - measurement.value
		case "down":
			position.y = position.y + measurement.value
		case "up":
			position.y = position.y - measurement.value
		}
	}
	log.Println("pt1", position.x*position.y)

	// pt 2
	var position2 Position

	for _, measurement := range measurements {
		switch measurement.direction {
		case "forward":
			position2.x = position2.x + measurement.value
			position2.y = position2.y + position2.aim*measurement.value
		case "backward":
			position2.x = position2.x - measurement.value
		case "up":
			position2.aim = position2.aim - measurement.value
		case "down":
			position2.aim = position2.aim + measurement.value
		}

	}

	log.Println("pos2", position2.x*position2.y)

}
