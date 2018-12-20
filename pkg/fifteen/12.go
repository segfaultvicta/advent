package fifteen

import (
	"encoding/json"
	"strconv"
)

func parseForIntegers(thing interface{}, filter string) (output int) {
outer:
	switch v := thing.(type) {
	case map[string]interface{}:
		// object case
		for _, child := range v {
			if child == filter {
				break outer
			}
		}
		for _, child := range v {
			output += parseForIntegers(child, filter)
		}
	case []interface{}:
		// array case
		for _, child := range v {
			output += parseForIntegers(child, filter)
		}
	case float64:
		output = int(v)
	default:
		return 0
	}
	return output
}

func day12sideA(lines []string) string {
	var bytes []byte

	for _, line := range lines {
		byteline := []byte(line)
		for _, thisbyte := range byteline {
			bytes = append(bytes, thisbyte)
		}
	}

	var unparsed interface{}
	json.Unmarshal(bytes, &unparsed)

	return strconv.Itoa(parseForIntegers(unparsed, ""))
}

func day12sideB(lines []string) string {
	var bytes []byte

	for _, line := range lines {
		byteline := []byte(line)
		for _, thisbyte := range byteline {
			bytes = append(bytes, thisbyte)
		}
	}

	var unparsed interface{}
	json.Unmarshal(bytes, &unparsed)

	return strconv.Itoa(parseForIntegers(unparsed, "red"))
}
