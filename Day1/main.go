package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	f, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(f)
	total := 0
	var inputArr []int

	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		inputArr = append(inputArr, num)
	}
	flag := 0
	var memoize []int
	for flag != 1 {

		for _, num := range inputArr {
			if flag != 1 {
				total += num
				for _, val := range memoize {
					if total == val {
						fmt.Print("First duplicate:")
						fmt.Println(total)
						flag = 1
					}
				}
				memoize = append(memoize, total)
			}
		}
	}
	fmt.Print("Total:")
	fmt.Println(total)
}
