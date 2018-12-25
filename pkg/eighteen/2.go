package eighteen

import (
	"fmt"
	"strings"

	util "github.com/segfaultvicta/advent/pkg"
)

func day2sideA(lines []string) string {
	twos := 0
	threes := 0

	for _, s := range lines {
		if hasCharacterTwice(s) {
			twos++
		}
		if hasCharacterThrice(s) {
			threes++
		}
	}
	return fmt.Sprintln("Checksum:", twos*threes)
}

func day2sideB(lines []string) string {
	for _, a := range lines {
		for _, b := range lines {
			if a != b && stringDifference(a, b) == 1 {
				fmt.Println(a)
				fmt.Println(b)
				return "Found it!"
			}
		}
	}
	return "something terrible has happened"
}

func stringDifference(a string, b string) (score int) {
	for idx, char := range a {
		if char != rune(b[idx]) {
			score++
		}
	}
	return
}

func hasCharacterTwice(line string) bool {
	for _, char := range util.Alphabet {
		if strings.Count(line, string(char)) == 2 {
			return true
		}
	}
	return false
}

func hasCharacterThrice(line string) bool {
	for _, char := range util.Alphabet {
		if strings.Count(line, string(char)) == 3 {
			return true
		}
	}
	return false
}
