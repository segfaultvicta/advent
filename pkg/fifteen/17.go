package fifteen

import (
	"fmt"
	"strconv"
)

// this function takes in an integer binary-encoding the container the sum of
// which we are trying to calculate, and a slice of integers containing the
// actual container mapped to by each bit position in the encoded integer, and
// returns the sum of container capacities encoded by the integer
func decodeAndSum(encoded int, mapping []int) int {

	//fmt.Printf("% x\n", buffer.Bytes())
	bin := fmt.Sprintf("%020b", encoded)
	sum := 0
	for i := 19; i >= 0; i-- {
		if bin[i] == 49 {
			sum += mapping[i]
		}
	}
	return sum
}

func decode(encoded int, mapping []int) []int {
	bin := fmt.Sprintf("%020b", encoded)
	var ret []int
	for i := 19; i >= 0; i-- {
		if bin[i] == 49 {
			ret = append(ret, mapping[i])
		}
	}
	return ret
}

func day17sideA(lines []string) string {
	cardinality := 1048575
	count := 0

	var containers []int
	for _, line := range lines {
		tmp, _ := strconv.Atoi(line)
		containers = append(containers, tmp)
	}

	for i := 0; i < cardinality; i++ {
		if decodeAndSum(i, containers) == 150 {
			count++
		}
	}

	return strconv.Itoa(count)
}

func day17sideB(lines []string) string {
	cardinality := 1048575
	count := 0

	var containers []int
	for _, line := range lines {
		tmp, _ := strconv.Atoi(line)
		containers = append(containers, tmp)
	}

	var options []int
	for i := 0; i < cardinality; i++ {
		if decodeAndSum(i, containers) == 150 {
			options = append(options, i)
		}
	}

	best := 99999
	var bestOptions []int
	for i := 0; i < len(options); i++ {
		min := len(decode(options[i], containers))
		if min <= best {
			best = min
			bestOptions = append(bestOptions, options[i])
		}
	}

	fmt.Println(bestOptions)
	for i := 0; i < len(bestOptions); i++ {
		fmt.Println(decode(bestOptions[i], containers))
	}

	return strconv.Itoa(count)
}
