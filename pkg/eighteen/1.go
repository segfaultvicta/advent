package eighteen

import (
	"fmt"
	"strconv"

	util "github.com/segfaultvicta/advent/pkg"
)

func day1sideA(lines []string) string {
	sum := 0
	for _, s := range lines {
		n, _ := strconv.Atoi(s)
		sum += n
	}
	return fmt.Sprintln(sum)
}

func day1sideB(lines []string) string {
	frequency := 0
	knownFrequencies := []int{0}

	changes := make([]int, len(lines))
	for i, s := range lines {
		n, _ := strconv.Atoi(s)
		changes[i] = n
	}

	for {
		for _, s := range changes {
			frequency += s
			if util.CheckIntInSlice(frequency, knownFrequencies) {
				return fmt.Sprintln("Found repeated frequency: ", frequency)
			}
			knownFrequencies = append(knownFrequencies, frequency)
		}
	}
}
