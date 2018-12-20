package fifteen

import (
	"fmt"
	"regexp"
	"strconv"
)

func day16sideA(lines []string) string {
	return "n/i"
}

type sue struct {
	num       int
	first     string
	numFirst  int
	second    string
	numSecond int
	third     string
	numThird  int
}

func day16sideB(lines []string) string {
	var sue int
	var f, s, t string
	var nf, ns, nt int
	re := regexp.MustCompile("Sue ([0-9]+): ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+), ([a-z]+): ([0-9]+)")

	for _, line := range lines {
		pieces := re.FindStringSubmatch(line) // i dont even care
		sue, _ = strconv.Atoi(pieces[1])
		f = pieces[2]
		nf, _ = strconv.Atoi(pieces[3])
		s = pieces[4]
		ns, _ = strconv.Atoi(pieces[5])
		t = pieces[6]
		nt, _ = strconv.Atoi(pieces[7])

		// cats, trees indicate there are greater than that many
		// pomeranians, goldfish indicate there are fewer than that many
		// children, samoyeds, akitas, vizslas, cars, perfumes are trustworthy

		fmt.Println(sue, f, nf, s, ns, t, nt)

		fmt.Println(sue, f, nf, s, ns, t, nt)
	}

	return "n/i"
}
