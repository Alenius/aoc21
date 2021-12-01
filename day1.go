package main

import (
	"log"
	"os"
	"strconv"
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
	lines := ReadNewlineSeparatedFile(wd + "/input/1.test.txt")

	log.Println("lines", readInput(lines))
}
