package fifteen

import (
	"fmt"
	"strconv"
)

func deliciosity(s, b, h, c int) int {
	//fmt.Println("------------------------------")
	//fmt.Println("S", s, "B", b, "H", h, "C", c)
	cpc := 2 * s
	dur := 5*b - c
	fla := -2*s + -3*b + 5*h
	tex := -1*h + 5*c
	cal := 3*s + 3*b + 8*h + 8*c
	tmp := cpc * dur * fla * tex
	if cpc < 1 || dur < 1 || fla < 1 || tex < 1 {
		//fmt.Println("zeroing out!")
		tmp = 0
	}
	if s+b+h+c != 100 {
		tmp = 0
	}
	if cal != 500 {
		tmp = 0
	}
	//fmt.Println("capacity", cpc, "durability", dur, "flavour", fla, "texture", tex, "by your powers combined", tmp)
	return tmp
}

func day15sideA(lines []string) string {
	// C =  2s
	// D =        5b      -  c
	// F = -2s + -3b + 5h
	// T =             -h + 5c
	// none of C, D, F, or T can be 0
	//
	// 0 < 5b - c -> c < 5b
	// 0 < 5c - h -> h < 5c  (h < 25b)

	best := 0

	for s := 0; s < 100; s++ {
		//fmt.Println(s)
		for b := 0; b < 100; b++ {
			//fmt.Println(b)
			//if s+b > 100 {
			//	break
			//}
			for c := 0; c < 100; c++ {
				//fmt.Println(c)
				//for c := 0; c < 5*b; c++ {
				//if s+b+c > 100 {
				//	break
				//}
				for h := 0; h < 100; h++ {
					//fmt.Println(h)
					//for h := 0; h < 5*c; h++ {
					//if s+b+c+h != 100 {
					//	break
					//}
					cookie := deliciosity(s, b, h, c)
					if cookie > best {
						fmt.Printf("S %d B %d H %d C %d ->%d\n", s, b, h, c, cookie)
						best = cookie
					}
					//panic("at the disco")
				}
			}
		}
	}

	return strconv.Itoa(best)
}

func day15sideB(lines []string) string {
	return "n/i"
}
