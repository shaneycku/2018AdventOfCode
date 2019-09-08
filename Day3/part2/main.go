package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type fabricPiece struct {
	IDs    []string
	claims int
}
type order struct {
	ID         string
	OffsetLeft int
	OffsetTop  int
	Width      int
	Length     int
}

func main() {
	const gridX = 2000
	const gridY = 2000
	grid := [gridX][gridY]fabricPiece{}

	inputArr := getInput()
	for _, j := range inputArr {
		for x := j.OffsetLeft; x <= (j.OffsetLeft+j.Width)-1; x++ {
			for y := j.OffsetTop; y <= (j.OffsetTop+j.Length)-1; y++ {
				grid[x][y].claims = grid[x][y].claims + 1
				grid[x][y].IDs = append(grid[x][y].IDs, j.ID)
			}
		}
	}
	noClaim := ""
	for _, claim := range inputArr {

		// flag := 0
		// for i := claim.OffsetLeft; i <= claim.OffsetLeft+claim.Width-1; i++ {
		// 	for j := claim.OffsetTop; j <= claim.OffsetTop+claim.Length-1; j++ {
		// 		// fmt.Println(grid[i][j])
		// 		if grid[i][j].claims != 1 {
		// 			flag = 1
		// 			break
		// 		}
		// 	}
		// 	if flag == 1 {
		// 		break
		// 	}
		// }
		// if flag == 0 {
		// 	noClaim = claim.ID
		// 	break
		// }

		answer, err := checkClaim(claim, grid)
		if err != nil {
			break
		}
		noClaim = answer
	}

	fmt.Print((fmt.Sprintf("\n\n**************\nID with no Claims: %v", noClaim)))
}
func checkClaim(claim order, grid [2000][2000]fabricPiece) (string, error) {
	for i := claim.OffsetLeft; i <= claim.OffsetLeft+claim.Width-1; i++ {
		for j := claim.OffsetTop; j <= claim.OffsetTop+claim.Length-1; j++ {
			// fmt.Println(grid[i][j])
			if grid[i][j].claims != 1 {
				return "", errors.New("Invalid")
			}
		}

	}
	fmt.Println(claim)
	return claim.ID, nil

}
func getInput() []order {
	f, _ := os.Open("../input.txt")
	scanner := bufio.NewScanner(f)
	var inputArr []order

	for scanner.Scan() {
		line := scanner.Text()
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
