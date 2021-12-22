package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func day7() {
	wd, _ := os.Getwd()
	values := ReadCommaSeparatedFile(wd + "/input/7.txt")

	var startingPositions []int
	for _, value := range values {
		intValue, _ := strconv.Atoi(value)
		startingPositions = append(startingPositions, intValue)
	}
	sort.Ints(startingPositions)

	median := getMedian(startingPositions)
	average := getAverage(startingPositions)
	fmt.Println(median)

	veryHighValue := 100000000000 // just take some high initial cost
	bestCostAndLevel := [2]int{veryHighValue, 0}
	span := 10
	for i := -span; i <= span; i++ {
		goalLevel := average + i

		// cost := calculateCost(startingPositions, goalLevel)
		cost := calculateExpensiveCost(startingPositions, goalLevel)

		bestCost := bestCostAndLevel[0]
		if cost < bestCost {
			fmt.Println("cost" + strconv.Itoa(cost))
			fmt.Println("level" + strconv.Itoa(goalLevel))
			bestCostAndLevel = [2]int{cost, goalLevel}
		}
	}

	fmt.Println("Best cost and level - ")
	fmt.Print(bestCostAndLevel)
}

func calculateCost(intSlice []int, goalLevel int) int {
	totalCost := 0
	for _, val := range intSlice {
		totalCost += absDiffInt(val, goalLevel)
	}

	return totalCost
}

func calculateExpensiveCost(intSlice []int, goalLevel int) int {
	totalCost := 0
	for _, val := range intSlice {
		diff := absDiffInt(val, goalLevel)
		totalCost += (diff * (diff + 1)) / 2 // triangular number
	}

	return totalCost
}
