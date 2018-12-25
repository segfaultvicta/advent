package eighteen

import (
	"fmt"
	"regexp"
	"strconv"
)

func day3sideA(lines []string) string {
	re := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	claims := make([][][]int, 1000)
	for i := range claims {
		claims[i] = make([][]int, 1000)
		for j := range claims[i] {
			claims[i][j] = make([]int, 0)
		}
	}

	for _, line := range lines {
		pieces := re.FindStringSubmatch(line)
		id, _ := strconv.Atoi(pieces[1])
		xLoc, _ := strconv.Atoi(pieces[2])
		yLoc, _ := strconv.Atoi(pieces[3])
		width, _ := strconv.Atoi(pieces[4])
		height, _ := strconv.Atoi(pieces[5])

		for x := xLoc; x < xLoc+width; x++ {
			for y := yLoc; y < yLoc+height; y++ {
				claims[x][y] = append(claims[x][y], id)
			}
		}
	}

	sumOverlaps := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			if len(claims[x][y]) > 1 {
				sumOverlaps++
			}
		}
	}

	return fmt.Sprintf("Found %d overlapping square inches of fabric\n", sumOverlaps)
}

func day3sideB(lines []string) string {
	re := regexp.MustCompile("#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)")
	claims := make([][][]int, 1000)

	for i := range claims {
		claims[i] = make([][]int, 1000)
		for j := range claims[i] {
			claims[i][j] = make([]int, 0)
		}
	}

	for _, line := range lines {
		pieces := re.FindStringSubmatch(line)
		id, _ := strconv.Atoi(pieces[1])
		xLoc, _ := strconv.Atoi(pieces[2])
		yLoc, _ := strconv.Atoi(pieces[3])
		width, _ := strconv.Atoi(pieces[4])
		height, _ := strconv.Atoi(pieces[5])

		for x := xLoc; x < xLoc+width; x++ {
			for y := yLoc; y < yLoc+height; y++ {
				claims[x][y] = append(claims[x][y], id)
			}
		}
	}

	for _, line := range lines {
		pieces := re.FindStringSubmatch(line)
		id, _ := strconv.Atoi(pieces[1])
		xLoc, _ := strconv.Atoi(pieces[2])
		yLoc, _ := strconv.Atoi(pieces[3])
		width, _ := strconv.Atoi(pieces[4])
		height, _ := strconv.Atoi(pieces[5])

		if checkUniqueClaim(xLoc, yLoc, width, height, claims) {
			return fmt.Sprintf("Found unique claim with ID #%d!\n", id)
		}
	}

	return "Found no unique claims; something terrible has happened."
}

func checkUniqueClaim(xLoc, yLoc, width, height int, claims [][][]int) bool {
	for x := xLoc; x < xLoc+width; x++ {
		for y := yLoc; y < yLoc+height; y++ {
			if len(claims[x][y]) != 1 {
				return false
			}
		}
	}

	return true
}
