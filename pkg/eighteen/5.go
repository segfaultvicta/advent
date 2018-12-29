package eighteen

import (
	"fmt"
	"strings"

	u "github.com/segfaultvicta/advent/pkg"
)

func day5sideA(lines []string) string {
	chars := strings.Split(lines[0], "")

	terminated := false
	for !terminated {
		chars, terminated = polymerReaction(chars)
	}

	return fmt.Sprintf("%d characters remain after all reactions have ceased.\n", len(chars))
}

func day5sideB(lines []string) string {
	for _, letter := range u.AlphabetLower {
		terminated := false
		chars := removeUpcaseAndDownAndSplit(lines[0], letter)
		for !terminated {
			chars, terminated = polymerReaction(chars)
		}
		if strings.Contains(lines[0], string(letter)) {
			fmt.Printf("Letter %s terminates with %d characters remaining.\n", string(letter), len(chars))
		}
	}
	return "look up and use mk. I eyeball"
}

func removeUpcaseAndDownAndSplit(input string, letter rune) []string {
	r := strings.Replace(input, string(letter), "", -1)
	r = strings.Replace(r, strings.ToUpper(string(letter)), "", -1)
	return strings.Split(r, "")
}

func polymerReaction(chars []string) ([]string, bool) {
	reactionResult := make([]string, 0)
	reactionHasOccurred := false
	for index := 0; index < len(chars); index++ {
		if index == len(chars)-1 {
			reactionResult = append(reactionResult, chars[index])
		} else if reactsWith(chars[index], chars[index+1]) {
			index++
			reactionHasOccurred = true
		} else {
			reactionResult = append(reactionResult, chars[index])
		}
	}
	return reactionResult, !reactionHasOccurred
}

func reactsWith(a, b string) bool {
	return strings.ToUpper(a) == strings.ToUpper(b) &&
		(strings.ToUpper(a) != a || strings.ToUpper(b) != b) &&
		(strings.ToLower(a) != a || strings.ToLower(b) != b)
}
