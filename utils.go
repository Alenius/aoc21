package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadNewlineSeparatedFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		new_line := scanner.Text()
		lines = append(lines, new_line)
	}

	return lines
}
