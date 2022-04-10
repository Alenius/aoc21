package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	t "github.com/alenius/aoctools"
)

type lanternfish struct {
	dayToSpawn int
}

func day6SolvePt1(startingLanternfish []lanternfish) {
	lanternfishes := append([]lanternfish{}, startingLanternfish...)

	numberOfDays := 100

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
		// fmt.Println("Day over with nr: " + strconv.Itoa(currentDay))
		currentDay++
	}

	log.Println(len(lanternfishes))
}

func day6SolvePt2(spawningSlice [9]int) {
	numberOfDays := 256

	currentDay := 1
	for currentDay <= numberOfDays {
		var nextDaySpawningSlice [9]int

		for i, value := range spawningSlice {
			if i == 0 {
				nextDaySpawningSlice[6] += value
				nextDaySpawningSlice[8] += value
			} else {
				nextDaySpawningSlice[i-1] += value
			}
		}

		currentDay++
		spawningSlice = nextDaySpawningSlice
	}

	totFishes := 0
	for _, totFishesPerDay := range spawningSlice {
		totFishes += totFishesPerDay
	}

	fmt.Println("No of fishes day 2  -  " + strconv.Itoa(totFishes))

}

func (aoc) Day6(inputPath string) {
	values := t.ReadCommaSeparatedFile(inputPath)
	var lanternfishes []lanternfish
	for _, value := range values {
		startingValue, _ := strconv.Atoi(value)
		fish := lanternfish{startingValue}
		lanternfishes = append(lanternfishes, fish)
	}

	startTime := time.Now()
	day6SolvePt1(lanternfishes)
	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)
	fmt.Println("Elapsed time pt1  -  " + elapsedTime.String())

	var spawningSlice [9]int
	for _, value := range values {
		startingValue, _ := strconv.Atoi(value)
		spawningSlice[startingValue] += 1
	}

	startTime = time.Now()
	day6SolvePt2(spawningSlice)
	endTime = time.Now()

	elapsedTime = endTime.Sub(startTime)
	fmt.Println("Elapsed time pt2  -  " + elapsedTime.String())

}
