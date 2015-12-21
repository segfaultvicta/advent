package main

import (
	"math"
	"strconv"
)

func DeduplicateIntegerSlice(a []int) []int {
	res := []int{}
	seen := map[int]int{}
	for _, val := range a {
		if _, ok := seen[val]; !ok {
			res = append(res, val)
			seen[val] = val
		}
	}
	return res
}

func FindDivisors(n int) []int {
	div := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			div = append(div, i)
		}
	}
	hdiv := []int{}
	for _, v := range div {
		if (n / v) != v {
			hdiv = append(hdiv, (n / v))
		}
	}
	for _, v := range hdiv {
		div = append(div, v)
	}

	div = DeduplicateIntegerSlice(div)
	return div
}

func FindLazyElfDivisors(n int) []int {
	div := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n)))+1; i++ {
		if n%i == 0 {
			div = append(div, i)
		}
	}
	hdiv := []int{}
	for _, v := range div {
		if (n / v) != v {
			hdiv = append(hdiv, (n / v))
		}
	}
	for _, v := range hdiv {
		div = append(div, v)
	}

	div = DeduplicateIntegerSlice(div)
	out := []int{}
	for _, v := range div {
		elfExhaustion := n / v
		if elfExhaustion <= 50 {
			out = append(out, v)
		}
	}
	// filter these divisors to those divisors which will have been applied 50 or fewer times
	return out
}

func day20sideA(lines []string) string {
	input := lines[0]

	test, _ := strconv.Atoi(input)
	found := false
	foundAt := 0
	for i := 1; found == false; i++ {
		//fmt.Println("House", i)
		div := FindDivisors(i)
		sum := 0
		for _, d := range div {
			sum += d * 10
		}
		//fmt.Println("gets", sum, "presents!")
		if sum >= test {
			found = true
			foundAt = i
		}
	}

	return strconv.Itoa(foundAt)
}

func day20sideB(lines []string) string {
	input := lines[0]

	test, _ := strconv.Atoi(input)
	found := false
	foundAt := 0
	for i := 1; found == false; i++ {
		//fmt.Println("House", i)
		div := FindLazyElfDivisors(i)
		sum := 0
		for _, d := range div {
			sum += d * 11
		}
		//fmt.Println("gets", sum, "presents!")
		if sum >= test {
			found = true
			foundAt = i
		}
	}

	return strconv.Itoa(foundAt)
}
