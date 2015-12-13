package main

import "strconv"

func indexToEndRun(s []byte, from int) int {
	for i := from + 1; i < len(s); i++ {
		if s[i] != s[from] {
			return i - 1
		}
	}
	return len(s) - 1
}

func day10sideA(lines []string) string {
	bytes := []byte(lines[0])
	for i := 0; i < 40; i++ {
		newString := []byte{}
		for j := 0; j < len(bytes); {
			runEndsAt := indexToEndRun(bytes, j)
			runSize := byte((runEndsAt + 1) - j)

			newString = append(newString, runSize+byte(48)) //shens
			newString = append(newString, bytes[j])

			j = runEndsAt + 1
		}
		bytes = newString
	}
	return strconv.Itoa(len(bytes))
}

func day10sideB(lines []string) string {
	bytes := []byte(lines[0])
	for i := 0; i < 50; i++ {
		newString := []byte{}
		for j := 0; j < len(bytes); {
			runEndsAt := indexToEndRun(bytes, j)
			runSize := byte((runEndsAt + 1) - j)

			newString = append(newString, runSize+byte(48)) //shens
			newString = append(newString, bytes[j])

			j = runEndsAt + 1
		}
		bytes = newString
	}
	return strconv.Itoa(len(bytes))
}
