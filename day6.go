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
			a_strs := strings.Split(split[1], ",")
			b_strs := strings.Split(split[3], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = !lights[j*1000+i]
				}
			}

		} else if split[1] == "on" {
			a_strs := strings.Split(split[2], ",")
			b_strs := strings.Split(split[4], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = true
				}
			}
		} else {
			a_strs := strings.Split(split[2], ",")
			b_strs := strings.Split(split[4], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = false
				}
			}
		}
	}
	count := 0
	for i := 0; i < 1000000; i++ {
		if lights[i] == true {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func day6sideB(lines []string) string {
	var lights [1000000]int
	for _, line := range lines {
		split := strings.Split(line, " ")

		if split[0] == "toggle" {
			a_strs := strings.Split(split[1], ",")
			b_strs := strings.Split(split[3], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = lights[j*1000+i] + 2
				}
			}

		} else if split[1] == "on" {
			a_strs := strings.Split(split[2], ",")
			b_strs := strings.Split(split[4], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = lights[j*1000+i] + 1
				}
			}
		} else {
			a_strs := strings.Split(split[2], ",")
			b_strs := strings.Split(split[4], ",")
			a_x, _ := strconv.Atoi(a_strs[0])
			a_y, _ := strconv.Atoi(a_strs[1])
			b_x, _ := strconv.Atoi(b_strs[0])
			b_y, _ := strconv.Atoi(b_strs[1])
			for i := a_x; i <= b_x; i++ {
				for j := a_y; j <= b_y; j++ {
					lights[j*1000+i] = lights[j*1000+i] - 1
					if lights[j*1000+i] < 0 {
						lights[j*1000+i] = 0
					}
				}
			}
		}
	}
	var count int = 0
	for i := 0; i < 1000000; i++ {
		count += lights[i]
	}
	return strconv.Itoa(count)
}
