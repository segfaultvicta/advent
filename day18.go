package main

import (
	"fmt"
	"strconv"
)

const XSize = 100
const YSize = 100
const XRange = XSize - 1
const YRange = YSize - 1
const XPenult = XRange - 1
const YPenult = YRange - 1

func PrintGame(game [YSize][XSize]bool) {
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

func CountOn(neighbors []bool) (ret int) {
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

func Count(game [YSize][XSize]bool) (ret int) {
	for _, row := range game {
		for _, cell := range row {
			if cell {
				ret++
			}
		}
	}
	return ret
}

func Evolve(game [YSize][XSize]bool) [YSize][XSize]bool {
	var ret [YSize][XSize]bool
	for y := 0; y < YSize; y++ {
		for x := 0; x < XSize; x++ {
			// how many neighbors of the current cell are active?
			var neighbors []bool
			jam := false
			if y == 0 || x == 0 || y == YRange || x == XRange {
				//we're on one of the edges and have to be careful
				switch {
				case y == 0 && x == 0:
					// upper left hand corner
					neighbors = append(neighbors, game[0][1])
					neighbors = append(neighbors, game[1][1])
					neighbors = append(neighbors, game[1][0])
					jam = true
				case y == 0 && x == XRange:
					// upper right hand corner
					neighbors = append(neighbors, game[0][XPenult])
					neighbors = append(neighbors, game[1][XPenult])
					neighbors = append(neighbors, game[1][XRange])
					jam = true
				case y == YRange && x == 0:
					// lower left hand corner
					neighbors = append(neighbors, game[YPenult][0])
					neighbors = append(neighbors, game[YPenult][1])
					neighbors = append(neighbors, game[YRange][1])
					jam = true
				case y == YRange && x == XRange:
					// lower right hand corner
					neighbors = append(neighbors, game[YPenult][XRange])
					neighbors = append(neighbors, game[YPenult][XPenult])
					neighbors = append(neighbors, game[YRange][XPenult])
					jam = true
				case y == 0:
					// upper edge
					neighbors = append(neighbors, game[0][x-1])
					neighbors = append(neighbors, game[0][x+1])
					neighbors = append(neighbors, game[1][x-1])
					neighbors = append(neighbors, game[1][x])
					neighbors = append(neighbors, game[1][x+1])
				case y == YRange:
					// lower edge
					neighbors = append(neighbors, game[YRange][x-1])
					neighbors = append(neighbors, game[YRange][x+1])
					neighbors = append(neighbors, game[YPenult][x-1])
					neighbors = append(neighbors, game[YPenult][x])
					neighbors = append(neighbors, game[YPenult][x+1])
				case x == 0:
					// left edge
					neighbors = append(neighbors, game[y+1][0])
					neighbors = append(neighbors, game[y-1][0])
					neighbors = append(neighbors, game[y+1][1])
					neighbors = append(neighbors, game[y][1])
					neighbors = append(neighbors, game[y-1][1])
				case x == XRange:
					// right edge
					neighbors = append(neighbors, game[y+1][XRange])
					neighbors = append(neighbors, game[y-1][XRange])
					neighbors = append(neighbors, game[y+1][XPenult])
					neighbors = append(neighbors, game[y][XPenult])
					neighbors = append(neighbors, game[y-1][XPenult])
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
			environment := CountOn(neighbors)
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
	var game [YSize][XSize]bool

	for y, line := range lines {
		for x, char := range line {
			if char == 35 {
				game[y][x] = true
			} else {
				game[y][x] = false
			}
		}
	}

	PrintGame(game)

	for i := 0; i < 100; i++ {
		game = Evolve(game)
		PrintGame(game)
	}

	return strconv.Itoa(Count(game))
}

func day18sideB(lines []string) string {
	count := 0

	return strconv.Itoa(count)
}
