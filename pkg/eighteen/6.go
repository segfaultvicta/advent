package eighteen

import (
	"fmt"
	"math"
	"strings"

	u "github.com/segfaultvicta/advent/pkg"
)

type vec struct {
	x int
	y int
}

const mapsize = 400
const saferange = 10000

func day6sideA(lines []string) string {
	m := make([][]int, mapsize)
	coords := make(map[int]vec)
	for x := 0; x < len(m); x++ {
		row := make([]int, mapsize)
		m[x] = row
	}

	coordID := 1
	for _, line := range lines {
		split := strings.Split(line, ", ")
		x := u.I(split[0])
		y := u.I(split[1])

		m[x][y] = coordID
		coords[coordID] = vec{x: x, y: y}
		coordID++
	}

	for x, row := range m {
		for y := range row {
			if m[x][y] == 0 {
				// find the nearest coordinate
				bestDistance := math.MaxInt32
				bestCoord := 0
				for id, coord := range coords {
					distance := u.Abs(coord.x-x) + u.Abs(coord.y-y)
					if distance < bestDistance {
						bestCoord = id
						bestDistance = distance
					} else if distance == bestDistance {
						bestCoord = -1
						bestDistance = distance
					}
				}
				m[x][y] = bestCoord
			}
		}
	}

	bannedCoordinates := make(map[int]bool)

	// eliminate any area that touches the edge of the map, b/c it's infinite
	for x, row := range m {
		for y := range row {
			if x == 0 || y == 0 || x == mapsize-1 || y == mapsize-1 {
				if !bannedCoordinates[m[x][y]] {
					bannedCoordinates[m[x][y]] = true
				}
			}
		}
	}

	for x, row := range m {
		for y := range row {
			if bannedCoordinates[m[x][y]] || m[x][y] == -1 {
				m[x][y] = 0
			}
		}
	}

	areas := make(map[int]int)

	for x, row := range m {
		for y := range row {
			if areas[m[x][y]] != 0 {
				areas[m[x][y]]++
			} else {
				areas[m[x][y]] = 1
			}
		}
	}

	/*for x, row := range m {
		for y := range row {
			if m[y][x] > 0 {
				fmt.Print(m[y][x])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}*/

	biggestArea := 0
	for id, area := range areas {
		if id != 0 && area > biggestArea {
			biggestArea = area
		}
	}

	return fmt.Sprintf("biggest area: %d\n", biggestArea)
}

func day6sideB(lines []string) string {
	m := make([][]int, mapsize)
	coords := make(map[int]vec)
	for x := 0; x < len(m); x++ {
		row := make([]int, mapsize)
		m[x] = row
	}

	coordID := 1
	for _, line := range lines {
		split := strings.Split(line, ", ")
		x := u.I(split[0])
		y := u.I(split[1])

		m[x][y] = coordID
		coords[coordID] = vec{x: x, y: y}
		coordID++
	}

	for x, row := range m {
		for y := range row {
			// find the sum-distance to all coordinates
			sumDistance := 0
			for _, coord := range coords {
				sumDistance += u.Abs(coord.x-x) + u.Abs(coord.y-y)
			}
			if sumDistance < saferange {
				m[x][y] = -1
			}
		}
	}

	/*for x, row := range m {
		for y := range row {
			if m[y][x] == -1 {
				fmt.Print("#")
			} else if m[y][x] == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(m[y][x])
			}
		}
		fmt.Print("\n")
	}*/

	safeCoords := 0
	for x, row := range m {
		for y := range row {
			if m[x][y] == -1 {
				safeCoords++
			}
		}
	}
	return fmt.Sprintf("Safe coordinates: %d", safeCoords)
}
