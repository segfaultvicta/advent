package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"
)

func main() {
	var day, side string
	flag.StringVar(&day, "day", "1", "index of puzzle we're solving")
	flag.StringVar(&side, "side", "A", "A for side A, B for side B")
	flag.Parse()
	side = strings.Trim(side, " ")
	day = strings.Trim(day, " ")

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

	if err := calendar[day+side](lines); err != nil {
		log.Fatal(err)
	}
}

var calendar = map[string]func([]string) error{
	"1A": day1sideA,
	"1B": day1sideB,
	"2A": day2sideA,
	"2B": day2sideB,
}
