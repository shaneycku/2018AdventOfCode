package main

import (
	"fmt"
	"strings"
)

func Part1() {
	double := 0
	triple := 0

	for _, val := range getInput() {
		letters := strings.Split(val, "")
		foundDouble := 0
		foundTriple := 0
		for _, letter := range letters {

			if strings.Count(val, letter) == 2 {
				foundDouble = 1
			}
			if strings.Count(val, letter) == 3 {
				foundTriple = 1
			}
		}
		if foundDouble == 1 {
			double++
			foundDouble = 0
		}
		if foundTriple == 1 {
			triple++
			foundTriple = 0
		}

	}

	fmt.Println("---------")
	fmt.Println(fmt.Sprintf("Double letters: %d", double))
	fmt.Println(fmt.Sprintf("Triple letters: %d", triple))
	fmt.Println(fmt.Sprintf("Checksum: %d", triple*double))

}

// func getInput() []string {
// 	f, _ := os.Open("./input.txt")
// 	scanner := bufio.NewScanner(f)
// 	var inputArr []string

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		// num, _ := strconv.Atoi(line)
// 		inputArr = append(inputArr, line)
// 	}
// 	return inputArr
// }
