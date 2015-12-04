package main

import (
	"bufio"
	"flag"
	"fmt"
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

		for _, e := range lines {
			pieces := strings.Split(e, "💩")
			fmt.Println("Testing side ", pieces[0], "of day ", day, "with input", pieces[1])
			// if test array ever contains newlines, this will have to change
			pieces_slice := []string{pieces[1]}
			result := calendar[day+pieces[0]](pieces_slice)
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
	"1A": day1sideA,
	"1B": day1sideB,
	"2A": day2sideA,
	"2B": day2sideB,
	"3A": day3sideA,
	"3B": day3sideB,
	"4A": day4sideA,
	"4B": day4sideB,
}
