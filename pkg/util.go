package util

import (
	"errors"
	"regexp"
	"strconv"
)

// Alphabet is a string containing the letters a through z
const Alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// FindIntInSlice looks for needle in haystack.
// If needle is found, returns the index of needle within haystack and nil.
// If needle isn't found, returns nil and an error indicating that needle wasn't found.
func FindIntInSlice(needle int, haystack []int) (index int, err error) {
	for i, e := range haystack {
		if e == needle {
			return i, nil
		}
	}
	return 0, errors.New("element was not found in slice")
}

// CheckIntInSlice returns true if needle is in haystack and false otherwise.
func CheckIntInSlice(needle int, haystack []int) (found bool) {
	_, err := FindIntInSlice(needle, haystack)
	return err == nil
}

// Parse parses line according to regexp and returns a map of named capture groups to captures.
func Parse(re *regexp.Regexp, line string) map[string]string {
	matches := re.FindStringSubmatch(line)

	params := make(map[string]string)
	for i, n := range re.SubexpNames() {
		if i > 0 && i < len(matches) && len(n) > 0 && len(matches[i]) > 0 {
			params[n] = matches[i]
		}
	}
	return params
}

// I returns the integer value of toInt assuming it'll parse correctly.
func I(toInt string) int {
	val, _ := strconv.Atoi(toInt)
	return val
}
