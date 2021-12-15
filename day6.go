package main

import (
	"log"
	"os"
	"strconv"
)

type lanternfish struct {
	dayToSpawn int
}

func day6() {
	wd, _ := os.Getwd()
	values := ReadCommaSeparatedFile(wd + "/input/6.txt")
	var lanternfishes []lanternfish
	for _, value := range values {
		startingValue, _ := strconv.Atoi(value)
		fish := lanternfish{startingValue}
		lanternfishes = append(lanternfishes, fish)
	}

	numberOfDays := 80

	currentDay := 1
	for currentDay <= numberOfDays {
		var newFishesForToday []lanternfish
		for i, fish := range lanternfishes {
			if fish.dayToSpawn == 0 {
				newFishesForToday = append(newFishesForToday, lanternfish{8})
				lanternfishes[i] = lanternfish{6}
			} else {
				lanternfishes[i] = lanternfish{fish.dayToSpawn - 1}
			}
		}

		lanternfishes = append(lanternfishes, newFishesForToday...)
		currentDay++
	}

	log.Println(len(lanternfishes))
}
