package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// var grid [][]int
	const gridX = 2000
	const gridY = 2000
	grid := [gridX][gridY]int{}

	inputArr := getInput
	for _, j := range inputArr() {

		// for x := gridX - j.OffsetLeft; x >= (gridX - j.OffsetLeft - j.Width); x-- {
		// 	for y := gridY - j.OffsetTop; y >= (gridY - j.OffsetTop - j.Length); y-- {
		// 		// fmt.Printf("X:%d ||Y:%d || Claim:%d\n", x, y, grid[x][y])
		// 		grid[x][y] = grid[x][y] + 1
		// 	}
		// }

		// fmt.Println(j)

		for x := j.OffsetLeft; x <= (j.OffsetLeft+j.Width)-1; x++ {
			for y := j.OffsetTop; y <= (j.OffsetTop+j.Length)-1; y++ {

				// fmt.Printf("X:%d ||Y:%d || Claim:%d\n", x, y, grid[x][y])
				grid[x][y] = grid[x][y] + 1
			}
		}
		// fmt.Print(fmt.Sprintf("Order#: %v\n", j.ID))
	}

	count := 0
	// fmt.Print(fmt.Sprintf("\n\n**************\nCount: %d", count))

	for ii := 1; ii <= gridX-1; ii++ {
		for jj := 1; jj <= gridY-1; jj++ {
			if grid[ii][jj] >= 2 {
				count++
			}
		}
	}
	fmt.Print((fmt.Sprintf("\n\n**************\nCount: %d", count)))
}

func getInput() []order {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)
	var inputArr []order

	for scanner.Scan() {
		line := scanner.Text()
		// num, _ := strconv.Atoi(line)
		re, _ := regexp.Compile(`(?m)#([\d]+)[\s]*@[\s]([\d]+),([\d]+):[\s]*([\d]+)x([\d]+)`)
		results := re.FindStringSubmatch(line)
		// fmt.Println(results)

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

type order struct {
	ID         string
	OffsetLeft int
	OffsetTop  int
	Width      int
	Length     int
}
