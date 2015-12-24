package main

import (
	"fmt"
	"math"
	"strconv"
)

func DecodeDay24(encoded int, mapping []int, magicFormatString string) []int {
	bin := fmt.Sprintf(magicFormatString, encoded) // still too lazy to do bit shifting
	var ret []int
	for i := len(mapping) - 1; i >= 0; i-- {
		if bin[i] == 49 {
			ret = append(ret, mapping[i])
		}
	}
	return ret
}

func day24sideA(lines []string) string {
	packages := []int{}
	for _, line := range lines {
		weight, _ := strconv.Atoi(line)
		packages = append(packages, weight)
	}

	fmt.Println(packages)

	sum := 0

	for _, pkg := range packages {
		sum += pkg
	}

	divsum := sum / 3

	//fmt.Println(sum, divsum)
	fmt.Printf("total weight %d, each division should weigh %d\n", sum, divsum)

	// subset-sum problem hurray! find all subsets that add up to divsum
	cardinality := int(math.Exp2(float64(len(packages))))
	//fmt.Println(cardinality, len(packages))
	fmt.Printf("%d packages, powerset cardinality %d\n", len(packages), cardinality)

	valid := [][]int{}

	bestLen := 6
	bestProd := 10439961859

	for subseti := 0; subseti < cardinality; subseti++ {
		if subseti%100000 == 0 {
			fmt.Println(bestLen, bestProd, subseti, len(valid), float64(subseti)/float64(cardinality))
		}
		subset := DecodeDay24(subseti, packages, "%028b")
		if len(subset) > bestLen {
			continue
		}
		subsetsum := 0
		subsetprod := 1
		for _, i := range subset {
			subsetsum += i
			subsetprod *= i
		}
		if subsetsum != divsum {
			continue
		}
		if subsetprod > bestProd {
			continue
		}
		// this isn't good enough - we need the set of subsets where the remainder can also be split
		// evenly into two sets of sum divsum.
		remainder := []int{}
		for _, i := range packages {
			found := false
			for _, j := range subset {
				if j == i {
					found = true
				}
			}
			if !found {
				remainder = append(remainder, i)
			}
		}
		// i'm going to hell
		magic := "%0" + strconv.Itoa(len(remainder)) + "b"
		remcardinality := int(math.Exp2(float64(len(remainder))))
		found := false
		for remsubseti := 0; (remsubseti < remcardinality) && !found; remsubseti++ {
			remsubset := DecodeDay24(remsubseti, remainder, magic)
			remsubsetsum := 0
			for _, i := range remsubset {
				remsubsetsum += i
			}
			if remsubsetsum != divsum {
				continue
			}
			if len(remsubset) == len(remainder) {
				continue
			}
			// i'm so sorry
			remremainder := []int{}
			for _, i := range remainder {
				foundtemp := false
				for _, j := range remsubset {
					if j == i {
						foundtemp = true
					}
				}
				if !foundtemp {
					remremainder = append(remremainder, i)
				}
			}
			remremaindersum := 0
			for _, i := range remremainder {
				remremaindersum += i
			}
			//fmt.Println("subset", subset, subsetsum, "remainder", remainder, "remsubset", remsubset, remsubsetsum, "remremainder", remremainder, remremaindersum)
			if remremaindersum != divsum || remsubsetsum != divsum {
				panic("remremainder or remainder are invalid, sleigh toppling nao")
			}
			found = true
		}
		if !found {
			continue
		}
		//fmt.Println(subset, subsetsum, remainder)
		valid = append(valid, subset)
		if len(subset) < bestLen {
			bestLen = len(subset)
		}
		if subsetprod < bestProd {
			bestProd = subsetprod
		}
	}

	fmt.Println(len(valid))

	return strconv.Itoa(bestProd)
}

func day24sideB(lines []string) string {
	packages := []int{}
	for _, line := range lines {
		weight, _ := strconv.Atoi(line)
		packages = append(packages, weight)
	}

	fmt.Println(packages)

	sum := 0

	for _, pkg := range packages {
		sum += pkg
	}

	divsum := sum / 4

	//fmt.Println(sum, divsum)
	fmt.Printf("total weight %d, each division should weigh %d\n", sum, divsum)

	// subset-sum problem hurray! find all subsets that add up to divsum
	cardinality := int(math.Exp2(float64(len(packages))))
	//fmt.Println(cardinality, len(packages))
	fmt.Printf("%d packages, powerset cardinality %d\n", len(packages), cardinality)

	valid := [][]int{}

	bestLen := 5
	bestProd := 72050269

	for subseti := 0; subseti < cardinality; subseti++ {
		if subseti%100000 == 0 {
			fmt.Println(bestLen, bestProd, subseti, len(valid), float64(subseti)/float64(cardinality))
		}
		subset := DecodeDay24(subseti, packages, "%028b")
		if len(subset) > bestLen {
			continue
		}
		subsetsum := 0
		subsetprod := 1
		for _, i := range subset {
			subsetsum += i
			subsetprod *= i
		}
		if subsetsum != divsum {
			continue
		}
		if subsetprod > bestProd {
			continue
		}
		// this isn't good enough - we need the set of subsets where the remainder can also be split
		// evenly into three! sets of sum divsum.
		remainder := []int{}
		for _, i := range packages {
			found := false
			for _, j := range subset {
				if j == i {
					found = true
				}
			}
			if !found {
				remainder = append(remainder, i)
			}
		}
		// i'm going to hell
		magic := "%0" + strconv.Itoa(len(remainder)) + "b"
		remcardinality := int(math.Exp2(float64(len(remainder))))
		found := false
		for remsubseti := 0; (remsubseti < remcardinality) && !found; remsubseti++ {
			remsubset := DecodeDay24(remsubseti, remainder, magic)
			remsubsetsum := 0
			for _, i := range remsubset {
				remsubsetsum += i
			}
			if remsubsetsum != divsum {
				continue
			}
			if len(remsubset) == len(remainder) {
				continue
			}
			// i'm so sorry
			remremainder := []int{}
			for _, i := range remainder {
				foundtemp := false
				for _, j := range remsubset {
					if j == i {
						foundtemp = true
					}
				}
				if !foundtemp {
					remremainder = append(remremainder, i)
				}
			}
			remremaindersum := 0
			for _, i := range remremainder {
				remremaindersum += i
			}
			// let's do the regrets warp again
			trunk := remremainder
			trunkmagic := "%0" + strconv.Itoa(len(trunk)) + "b"
			trunkcardinality := int(math.Exp2(float64(len(trunk))))
			foundtrunk := false
			for trunksubseti := 0; (trunksubseti < trunkcardinality) && !foundtrunk; trunksubseti++ {
				trunksubset := DecodeDay24(trunksubseti, trunk, trunkmagic)
				trunksubsetsum := 0
				for _, i := range trunksubset {
					trunksubsetsum += i
				}
				if trunksubsetsum != divsum {
					continue
				}
				if len(trunksubset) == len(trunk) {
					continue
				}
				trunkremainder := []int{}
				for _, i := range trunk {
					foundtemp := false
					for _, j := range trunksubset {
						if j == i {
							foundtemp = true
						}
					}
					if !foundtemp {
						trunkremainder = append(trunkremainder, i)
					}
				}
				trunkremaindersum := 0
				for _, i := range trunkremainder {
					trunkremaindersum += i
				}

				if remsubsetsum != divsum || trunksubsetsum != divsum || trunkremaindersum != divsum {
					//fmt.Println("subset", subset, subsetsum, "remainder", remainder, "remsubset", remsubset, remsubsetsum, "remremainder", remremainder, remremaindersum)
					fmt.Println("subset", subset, "subset sum", subsetsum, "remainder", remainder)
					fmt.Println("remsubset", remsubset, "remsubset sum", remsubsetsum, "remremainder", remremainder)
					fmt.Println("trunksubsetsum", trunksubsetsum, "trunkremaindersum", trunkremaindersum)

					panic("remremainder or remsubsetsum or trunksum are invalid, sleigh toppling nao")
				}
				foundtrunk = true
			}
			if !foundtrunk {
				continue
			}
			found = true
		}
		if !found {
			continue
		}
		//fmt.Println(subset, subsetsum, remainder)
		valid = append(valid, subset)
		if len(subset) < bestLen {
			bestLen = len(subset)
		}
		if subsetprod < bestProd {
			bestProd = subsetprod
		}
	}

	fmt.Println(len(valid))

	return strconv.Itoa(bestProd)
}
