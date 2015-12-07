package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("(?P<a>[0-9a-z]*)? ?(?P<op>AND|OR|NOT|LSHIFT|RSHIFT)? ?(?P<b>[0-9a-z]*) -> (?P<out>[a-z]*)")

func day7sideA(lines []string) string {
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		r := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 {
				r[name] = match[i]
			}
		}

		fmt.Println("parsing", line)
		p := r["a"]
		q := r["b"]
		d := r["out"]
		switch r["op"] {
		case "AND":
		case "OR":
		case "LSHIFT":
		case "RSHIFT":
		case "NOT":
		default:
			// assignment of input to a wire
		}
	}

	return "n/i"
}

func day7sideB(lines []string) string {
	return "n/i"
}
