package fifteen

import (
	"fmt"
	"regexp"
	"strconv"
)

func day8sideA(lines []string) string {
	stripUnicode := regexp.MustCompile(`\\x[0-9a-fA-F][0-9a-fA-F]`)
	stripQuot := regexp.MustCompile(`\\\"`)
	stripBw := regexp.MustCompile(`\\\\`)
	count := 0
	for _, line := range lines {
		count = count + len(line)
		line = stripUnicode.ReplaceAllString(line, "*")
		line = stripQuot.ReplaceAllString(line, "\"")
		line = stripBw.ReplaceAllString(line, "\\")
		line = line[1 : len(line)-1]
		count = count - len(line)
	}
	return strconv.Itoa(count)
}

func day8sideB(lines []string) string {
	count := 0
	for _, line := range lines {
		lineCount := len(line)

		fmt.Println("--------------------------------")
		fmt.Println(line)

		var build []byte
		for i := 0; i < len(line); i++ {
			if line[i] == 92 || line[i] == 34 {
				build = append(build, 92)
				build = append(build, line[i])
			} else {
				build = append(build, line[i])
			}
		}

		line = "\"" + (string(build)) + "\""

		fmt.Println(line)

		lineCount = len(line) - lineCount
		count = count + lineCount
	}
	return strconv.Itoa(count)
}
