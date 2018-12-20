package fifteen

import (
	"sort"
	"strconv"
	"strings"
)

func day2sideA(lines []string) string {
	count := 0
	for _, e := range lines {
		d := splitline(e)
		s := []int{a(d, 0, 1), a(d, 1, 2), a(d, 2, 0)}
		sort.Ints(s)
		count += 2*s[0] + 2*s[1] + 2*s[2] + s[0]
	}
	return strconv.Itoa(count)
}

func day2sideB(lines []string) string {
	count := 0
	for _, e := range lines {
		d := splitline(e)
		s := []int{p(d, 0, 1), p(d, 1, 2), p(d, 2, 0)}
		sort.Ints(s)
		count += d[0]*d[1]*d[2] + s[0]
	}
	return strconv.Itoa(count)
}

func splitline(line string) [3]int {
	splitstr := strings.Split(line, "x")
	var dimensions [3]int
	dimensions[0], _ = strconv.Atoi(splitstr[0])
	dimensions[1], _ = strconv.Atoi(splitstr[1])
	dimensions[2], _ = strconv.Atoi(splitstr[2])

	return dimensions
}

func p(dims [3]int, i int, j int) int {
	return dims[i]*2 + dims[j]*2
}

func a(dims [3]int, i int, j int) int {
	return dims[i] * dims[j]
}
