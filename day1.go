package main

import (
	"log"
	"os"
	"strconv"

	t "github.com/alenius/aoctools"
)

type Measurement struct {
	value                         int
	relationToPreviousMeasurement string
}

func (Measurement) create(value int, relation string) Measurement {
	return Measurement{value: value, relationToPreviousMeasurement: relation}
}

func readInput(lines []string) []Measurement {
	var measurements []Measurement
	for i, s := range lines {
		destringifiedValue, _ := strconv.Atoi(s)
		if i == 0 {
			new_measurement := Measurement{}.create(destringifiedValue, "n/a")
			measurements = append(measurements, new_measurement)
		} else {
			prev_value := measurements[len(measurements)-1]
			var relation string

			switch rel := destringifiedValue - prev_value.value; {
			case (rel > 0):
				relation = "+"
			case (rel == 0):
				relation = "/"
			case (rel < 0):
				relation = "-"
			}

			new_measurement := Measurement{}.create(destringifiedValue, relation)
			measurements = append(measurements, new_measurement)
		}
	}

	return measurements
}

func day1() {
	wd, _ := os.Getwd()
	lines := t.ReadNewlineSeparatedFile(wd + "/input/1.txt")

	measurements := readInput(lines)

	noOfLarger := 0

	// pt1
	for _, m := range measurements {
		if m.relationToPreviousMeasurement == "+" {
			noOfLarger++
		}
	}

	log.Println("number of increasing", noOfLarger)

	// pt2
	greatestSumRun := 0
	previousSum := 0
	for i, m := range measurements {
		if i == 0 || i == 1 {
			continue
		}

		newSum := m.value + measurements[i-1].value + measurements[i-2].value
		if newSum > previousSum && previousSum != 0 {
			greatestSumRun++
			log.Println("increased")

		}

		previousSum = newSum
	}

	log.Println("greatest sum run", greatestSumRun)

}
