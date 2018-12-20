package fifteen

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

func day4sideA(lines []string) string {
	for i := 0; i < 1000000; i++ {
		h := md5.New()
		in := lines[0] + strconv.Itoa(i)
		io.WriteString(h, in)
		first5 := fmt.Sprintf("%x", h.Sum(nil))[:5]
		if first5 == "00000" {
			return strconv.Itoa(i)
		}
	}
	return "No match found"
}

func day4sideB(lines []string) string {
	for i := 0; i < 10000000; i++ {
		h := md5.New()
		in := lines[0] + strconv.Itoa(i)
		io.WriteString(h, in)
		first6 := fmt.Sprintf("%x", h.Sum(nil))[:6]
		if first6 == "000000" {
			return strconv.Itoa(i)
		}
	}
	return "No match found"
}
