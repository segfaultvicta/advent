package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day5sideA(lines []string) string {
	nicecount := 0
	for _, line := range lines {
		vowel, _ := regexp.MatchString("[aeiou].*[aeiou].*[aeiou].*", line)
		double := doubledLetter(line)
		snowflake, _ := regexp.MatchString("ab|cd|pq|xy", line)
		if vowel && double && !snowflake {
			nicecount += 1
		}
	}
	return strconv.Itoa(nicecount)
}

func doubledLetter(line string) bool {
	for i := 1; i < len(line); i++ {
		if line[i] == line[i-1] {
			return true
		}
	}
	return false
}

func suckit(line string) bool {
	//foreach pair of letters in the string
	for i := 0; i < len(line)-1; i++ {
		if strings.Contains(line[i+2:], line[i:i+2]) {
			return true
		}
	}
	return false
}

func sandwich(line string) bool {
	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}
	return false
}

func day5sideB(lines []string) string {
	nicecount := 0
	for _, line := range lines {
		suck := suckit(line)
		sand := sandwich(line)
		if suck && sand {
			nicecount += 1
		}
	}
	return strconv.Itoa(nicecount)
}
