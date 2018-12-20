package main

import (
	"fmt"
	"strconv"
)

const xSize = 100
const ySize = 100
const xRange = xSize - 1
const yRange = ySize - 1
const xPenult = xRange - 1
const yPenult = yRange - 1

func printGame(game [ySize][xSize]bool) {
	ret := ""
	for _, row := range game {
		str := ""
		for _, set := range row {
			if set {
				str += "#"
			} else {
				str += "."
			}

		}
		ret += str + "\n"
	}
	fmt.Println(ret)
}

func countOn(neighbors []bool) (ret int) {
	for _, neighbor := range neighbors {
		if neighbor {
			ret++
		}
	}
	//if ret > 0 {
	//	fmt.Println(neighbors, ret)
	//}
	return ret
}

func count(game [ySize][xSize]bool) (ret int) {
	for _, row := range game {
		for _, cell := range row {
			if cell {
				ret++
			}
		}
	}
	return ret
}

func evolve(game [ySize][xSize]bool) [ySize][xSize]bool {
	var ret [ySize][xSize]bool
	for y := 0; y < ySize; y++ {
		for x := 0; x < xSize; x++ {
			// how many neighbors of the current cell are active?
			var neighbors []bool
			jam := false
			if y == 0 || x == 0 || y == yRange || x == xRange {
				//we're on one of the edges and have to be careful
				switch {
				case y == 0 && x == 0:
					// upper left hand corner
					neighbors = append(neighbors, game[0][1])
					neighbors = append(neighbors, game[1][1])
					neighbors = append(neighbors, game[1][0])
					jam = true
				case y == 0 && x == xRange:
					// upper right hand corner
					neighbors = append(neighbors, game[0][xPenult])
					neighbors = append(neighbors, game[1][xPenult])
					neighbors = append(neighbors, game[1][xRange])
					jam = true
				case y == yRange && x == 0:
					// lower left hand corner
					neighbors = append(neighbors, game[yPenult][0])
					neighbors = append(neighbors, game[yPenult][1])
					neighbors = append(neighbors, game[yRange][1])
					jam = true
				case y == yRange && x == xRange:
					// lower right hand corner
					neighbors = append(neighbors, game[yPenult][xRange])
					neighbors = append(neighbors, game[yPenult][xPenult])
					neighbors = append(neighbors, game[yRange][xPenult])
					jam = true
				case y == 0:
					// upper edge
					neighbors = append(neighbors, game[0][x-1])
					neighbors = append(neighbors, game[0][x+1])
					neighbors = append(neighbors, game[1][x-1])
					neighbors = append(neighbors, game[1][x])
					neighbors = append(neighbors, game[1][x+1])
				case y == yRange:
					// lower edge
					neighbors = append(neighbors, game[yRange][x-1])
					neighbors = append(neighbors, game[yRange][x+1])
					neighbors = append(neighbors, game[yPenult][x-1])
					neighbors = append(neighbors, game[yPenult][x])
					neighbors = append(neighbors, game[yPenult][x+1])
				case x == 0:
					// left edge
					neighbors = append(neighbors, game[y+1][0])
					neighbors = append(neighbors, game[y-1][0])
					neighbors = append(neighbors, game[y+1][1])
					neighbors = append(neighbors, game[y][1])
					neighbors = append(neighbors, game[y-1][1])
				case x == xRange:
					// right edge
					neighbors = append(neighbors, game[y+1][xRange])
					neighbors = append(neighbors, game[y-1][xRange])
					neighbors = append(neighbors, game[y+1][xPenult])
					neighbors = append(neighbors, game[y][xPenult])
					neighbors = append(neighbors, game[y-1][xPenult])
				}
			} else {
				//we're inside and can be less careful
				neighbors = append(neighbors, game[y-1][x-1])
				neighbors = append(neighbors, game[y-1][x])
				neighbors = append(neighbors, game[y-1][x+1])
				neighbors = append(neighbors, game[y][x+1])
				neighbors = append(neighbors, game[y][x-1])
				neighbors = append(neighbors, game[y+1][x-1])
				neighbors = append(neighbors, game[y+1][x])
				neighbors = append(neighbors, game[y+1][x+1])
			}
			environment := countOn(neighbors)
			if game[y][x] {
				//fmt.Println("eep", environment)
				if environment == 2 || environment == 3 || jam {
					ret[y][x] = true
					//	fmt.Println("yeep", y, x)
				} else {
					ret[y][x] = false
					//	fmt.Println("weep", y, x)
				}
			} else {
				if environment == 3 || jam {
					//	fmt.Println("woop", y, x)
					ret[y][x] = true
				} else {
					ret[y][x] = false
				}
			}
		}
	}
	return ret
}

func day18sideA(lines []string) string {
	var game [ySize][xSize]bool

	for y, line := range lines {
		for x, char := range line {
			if char == 35 {
				game[y][x] = true
			} else {
				game[y][x] = false
			}
		}
	}

	printGame(game)

	for i := 0; i < 100; i++ {
		game = evolve(game)
		printGame(game)
	}

	return strconv.Itoa(count(game))
}

func day18sideB(lines []string) string {
	count := 0

	return strconv.Itoa(count)
}
