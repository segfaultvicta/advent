package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var day, side, test string
	flag.StringVar(&day, "day", "1", "index of puzzle we're solving")
	flag.StringVar(&side, "side", "A", "A for side A, B for side B")
	flag.StringVar(&test, "test", "N", "Are we running test instead of input? Y/N")
	flag.Parse()
	side = strings.Trim(side, " ")
	day = strings.Trim(day, " ")
	test = strings.Trim(test, " ")

	if test == "Y" {
		s := []string{"day", day, ".test"}
		infile := strings.Join(s, "")

		b, _ := ioutil.ReadFile(infile)
		contents := string(b)

		lines := strings.Split(contents, "ᚼ")

		for _, e := range lines {
			pieces := strings.Split(e, "ᛥ")
			fmt.Println("Testing side", pieces[0], "of day", day, "with input", pieces[1])
			split_by_newline := strings.Split(pieces[1], "\r\n")
			result := calendar[day+pieces[0]](split_by_newline)
			if pieces[2] == result {
				fmt.Println(".")
			} else {
				fmt.Println("Expected", pieces[2], "but got", result)
			}
		}
	} else {
		s := []string{"day", day, ".input"}
		infile := strings.Join(s, "")

		file, err := os.Open(infile)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(calendar[day+side](lines))
	}
}

var calendar = map[string]func([]string) string{
	"1A":  day1sideA,
	"1B":  day1sideB,
	"2A":  day2sideA,
	"2B":  day2sideB,
	"3A":  day3sideA,
	"3B":  day3sideB,
	"4A":  day4sideA,
	"4B":  day4sideB,
	"5A":  day5sideA,
	"5B":  day5sideB,
	"6A":  day6sideA,
	"6B":  day6sideB,
	"7A":  day7sideA,
	"7B":  day7sideB,
	"8A":  day8sideA,
	"8B":  day8sideB,
	"9A":  day9sideA,
	"9B":  day9sideB,
	"11A": day11sideA,
	"11B": day11sideB,
	"12A": day12sideA,
	"12B": day12sideB,
	"13A": day13sideA,
	"13B": day13sideB,
	"14A": day14sideA,
	"14B": day14sideB,
}
