package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main() {

}

type shiftEvent struct {
	date    string
	month   int
	day     int
	year    int
	hour    int
	minute  int
	action  string
	guardID string
}

func getInput() []order {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)
	var inputArr []order
	

	for scanner.Scan() {
		line := scanner.Text()
		var re = regexp.MustCompile(`(?m)\[(\d{4})-(\d{2})-(\d{2})[\s](\d{2}):(\d{2})\]\s+([\w\s]+(#(\d{4})*)*[\w\s]+)`)
		re, _ := regexp.Compile(`(?m)#([\d]+)[\s]*@[\s]([\d]+),([\d]+):[\s]*([\d]+)x([\d]+)`)
		results := re.FindStringSubmatch(line)

		offsetLeft, _ := strconv.Atoi(results[2])
		offsetTop, _ := strconv.Atoi(results[3])
		width, _ := strconv.Atoi(results[4])
		height, _ := strconv.Atoi(results[5])

		order := order{
			results[1], offsetLeft, offsetTop, width, height}
		inputArr = append(inputArr, order)
	}
	return inputArr
}
