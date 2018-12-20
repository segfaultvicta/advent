package main

import (
	"strconv"
	"strings"
)

func day6sideA(lines []string) string {
	var lights [1000000]bool
	for _, line := range lines {
		split := strings.Split(line, " ")

		if split[0] == "toggle" {
			aStrs := strings.Split(split[1], ",")
			bStrs := strings.Split(split[3], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = !lights[j*1000+i]
				}
			}

		} else if split[1] == "on" {
			aStrs := strings.Split(split[2], ",")
			bStrs := strings.Split(split[4], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = true
				}
			}
		} else {
			aStrs := strings.Split(split[2], ",")
			bStrs := strings.Split(split[4], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = false
				}
			}
		}
	}
	count := 0
	for i := 0; i < 1000000; i++ {
		if lights[i] == true {
			count++
		}
	}
	return strconv.Itoa(count)
}

func day6sideB(lines []string) string {
	var lights [1000000]int
	for _, line := range lines {
		split := strings.Split(line, " ")

		if split[0] == "toggle" {
			aStrs := strings.Split(split[1], ",")
			bStrs := strings.Split(split[3], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = lights[j*1000+i] + 2
				}
			}

		} else if split[1] == "on" {
			aStrs := strings.Split(split[2], ",")
			bStrs := strings.Split(split[4], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = lights[j*1000+i] + 1
				}
			}
		} else {
			aStrs := strings.Split(split[2], ",")
			bStrs := strings.Split(split[4], ",")
			aX, _ := strconv.Atoi(aStrs[0])
			aY, _ := strconv.Atoi(aStrs[1])
			bX, _ := strconv.Atoi(bStrs[0])
			bY, _ := strconv.Atoi(bStrs[1])
			for i := aX; i <= bX; i++ {
				for j := aY; j <= bY; j++ {
					lights[j*1000+i] = lights[j*1000+i] - 1
					if lights[j*1000+i] < 0 {
						lights[j*1000+i] = 0
					}
				}
			}
		}
	}
	var count int
	for i := 0; i < 1000000; i++ {
		count += lights[i]
	}
	return strconv.Itoa(count)
}
