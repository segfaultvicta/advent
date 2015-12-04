package main

import (
	"strconv"
)

// ^>v< 94 62 118 60

type Position struct {
	X int
	Y int
}

func NewPosition(x int, y int) *Position {
	var ret Position
	ret.X = x
	ret.Y = y
	return &ret
}

func day3sideA(lines []string) string {
	line := lines[0]
	var grid []Position
	curr := NewPosition(0, 0)
	grid = append(grid, *curr)
	count := 1
	for i := 0; i < len(line); i++ {
		next := nextPosition(curr, line[i])
		if !contains(grid, *next) {
			count += 1
			grid = append(grid, *next)
		}
		curr = next
	}
	return strconv.Itoa(count)
}

func day3sideB(lines []string) string {
	line := lines[0]
	var grid []Position
	currSanta := NewPosition(0, 0)
	currRobot := NewPosition(0, 0)
	roboturn := false
	grid = append(grid, *currSanta)
	count := 1
	for i := 0; i < len(line); i++ {
		if roboturn {
			// robo-santa!
			next := nextPosition(currRobot, line[i])
			if !contains(grid, *next) {
				count += 1
				grid = append(grid, *next)
			}
			roboturn = false
			currRobot = next
		} else {
			// actual santa!
			next := nextPosition(currSanta, line[i])
			if !contains(grid, *next) {
				count += 1
				grid = append(grid, *next)
			}
			roboturn = true
			currSanta = next
		}
	}
	return strconv.Itoa(count)
}

func nextPosition(curr *Position, char byte) *Position {
	var next *Position
	switch char {
	case 94:
		// up
		next = NewPosition(curr.X, curr.Y+1)
	case 62:
		// right
		next = NewPosition(curr.X+1, curr.Y)
	case 118:
		//down
		next = NewPosition(curr.X, curr.Y-1)
	case 60:
		//left
		next = NewPosition(curr.X-1, curr.Y)
	}
	return next
}

func contains(grid []Position, location Position) bool {
	for _, e := range grid {
		if location == e {
			return true
		}
	}
	return false
}
