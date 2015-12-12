package main

import (
	"encoding/json"
	"strconv"
)

type JsonObject map[string]interface{}

func parseForIntegers(thing interface{}, filter string) (int, bool) {
	sum := 0

	switch thing.(type) {
	case map[string]interface{}:
		// object case
		jsonObject := thing.(map[string]interface{})
		latch := false
		for _, child := range jsonObject {
			res, filtered := parseForIntegers(child, filter)
			if filtered {
				latch = true
				sum = 0
			}
			if !filtered && !latch {
				sum += res
			}
		}
	case []interface{}:
		// array case
		array := thing.([]interface{})
		for _, child := range array {
			res, _ := parseForIntegers(child, filter)
			sum += res
		}
	case string:
		node := thing.(string)
		if len(filter) > 0 && node == filter {
			return 0, true
		} else {
			return 0, false
		}
	case float64:
		node := thing.(float64)
		sum = int(node)
	default:
		return 0, false
	}
	return sum, false
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

	sum, _ := parseForIntegers(unparsed, "")
	return strconv.Itoa(sum)
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

	sum, _ := parseForIntegers(unparsed, "red")
	return strconv.Itoa(sum)
}
