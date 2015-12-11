package main

import (
	"regexp"
	"sort"
	"strings"
)

func smallestNonNegative(slice []int) int {
	sort.Ints(slice)
	for i := 0; i < len(slice); i++ {
		if slice[i] != -1 {
			return slice[i]
		}
	}
	return -1
}

func incrementPassword(password string, place int) string {
	matched, _ := regexp.MatchString("i|o|l", password)
	if matched {
		// god this is ugly
		mostSignificantI := strings.Index(password, "i")
		mostSignificantO := strings.Index(password, "o")
		mostSignificantL := strings.Index(password, "l")
		var indices = []int{mostSignificantI, mostSignificantO, mostSignificantL}
		index := smallestNonNegative(indices)

		ret := password[:index]
		cutpoint := password[index]
		ret += string(cutpoint + 1)
		ret += strings.Repeat("a", 7-index)

		return ret
	}
	last := password[place]
	ret := password[:place]
	//fmt.Println(last)
	//fmt.Println(ret)
	if last == 122 {
		last = 97
		ret = incrementPassword(ret, place-1)
	} else {
		last += 1
	}
	ret += string(last)
	return ret
}

func isValidPassword(password string) bool {
	//fmt.Println("checking", password, "for validity")
	matched, _ := regexp.MatchString("i|o|l", password)
	if matched {
		return false
	}
	pairs := 0

	for i := 0; i < 7; i++ {
		if password[i:i+1] == password[i+1:i+2] {
			pairs += 1
			i++
		}
	}
	if pairs < 2 {
		return false
	}

	for i := 0; i < 6; i++ {

		a := password[i]
		b := password[i+1]
		c := password[i+2]
		if (a+1 == b) && (b+1 == c) {
			return true
		}
	}
	return false
}

func day11sideA(lines []string) string {
	//password := lines[0]
	password := "hxbxxyzz"
	valid := false
	for !valid {
		password = incrementPassword(password, 7)
		valid = isValidPassword(password)
	}
	return password
}

func day11sideB(lines []string) string {
	return "n/i"
}
