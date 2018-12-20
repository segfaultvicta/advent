package main

import (
	"strconv"
)

// ^>v< 94 62 118 60

type position struct {
	X int
	Y int
}

func newPosition(x int, y int) *position {
	var ret position
	ret.X = x
	ret.Y = y
	return &ret
}

func day3sideA(lines []string) string {
	line := lines[0]
	var grid []position
	curr := newPosition(0, 0)
	grid = append(grid, *curr)
	count := 1
	for i := 0; i < len(line); i++ {
		next := nextPosition(curr, line[i])
		if !contains(grid, *next) {
			count++
			grid = append(grid, *next)
		}
		curr = next
	}
	return strconv.Itoa(count)
}

func day3sideB(lines []string) string {
	line := lines[0]
	var grid []position
	currSanta := newPosition(0, 0)
	currRobot := newPosition(0, 0)
	roboturn := false
	grid = append(grid, *currSanta)
	count := 1
	for i := 0; i < len(line); i++ {
		if roboturn {
			// robo-santa!
			next := nextPosition(currRobot, line[i])
			if !contains(grid, *next) {
				count++
				grid = append(grid, *next)
			}
			roboturn = false
			currRobot = next
		} else {
			// actual santa!
			next := nextPosition(currSanta, line[i])
			if !contains(grid, *next) {
				count++
				grid = append(grid, *next)
			}
			roboturn = true
			currSanta = next
		}
	}
	return strconv.Itoa(count)
}

func nextPosition(curr *position, char byte) *position {
	var next *position
	switch char {
	case 94:
		// up
		next = newPosition(curr.X, curr.Y+1)
	case 62:
		// right
		next = newPosition(curr.X+1, curr.Y)
	case 118:
		//down
		next = newPosition(curr.X, curr.Y-1)
	case 60:
		//left
		next = newPosition(curr.X-1, curr.Y)
	}
	return next
}

func contains(grid []position, location position) bool {
	for _, e := range grid {
		if location == e {
			return true
		}
	}
	return false
}
