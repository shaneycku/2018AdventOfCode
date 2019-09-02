package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// var grid [][]int
	grid := [1600][1600]int{}
	// grid[0][100] = 10
	re := regexp.MustCompile(`@[\s]([0-9]+),[0-9]+:`)

	fmt.Print(grid[0][100])
}

func getInput() []string {
	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	var inputArr []string

	for scanner.Scan() {
		line := scanner.Text()
		// num, _ := strconv.Atoi(line)
		inputArr = append(inputArr, line)
	}
	return inputArr
}
