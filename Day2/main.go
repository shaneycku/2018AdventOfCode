package main

import (
	"bufio"
	"os"
)

func main() {
	Part1()
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
