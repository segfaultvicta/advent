package main

import "strconv"

func day25sideA(lines []string) string {
	row := 2947
	//row := 3
	column := 3029
	//column := 3

	yIncrement := 0
	y := 1

	var nthnumber int

	for i := 1; i <= row; i++ {
		y = y + yIncrement

		//fmt.Printf("% 4d ", i, y)
		yIncrement++

		xIncrement := 1 + yIncrement
		x := y
		for j := 2; j <= column; j++ {
			x = x + xIncrement
			//fmt.Printf("% 4d ", x)
			if i == row && j == column {
				nthnumber = x
			}
			xIncrement++
		}
	}

	magicnumber := 20151125
	//nthnumber = 10
	//fmt.Println(magicnumber)
	for i := 1; i < nthnumber; i++ {
		magicnumber *= 252533
		magicnumber = magicnumber % 33554393
		//fmt.Println(magicnumber)
	}

	return strconv.Itoa(magicnumber)
}

func day25sideB(lines []string) string {
	return "n/i"
}
