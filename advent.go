package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/segfaultvicta/advent/pkg/eighteen"
	"github.com/segfaultvicta/advent/pkg/fifteen"
	"github.com/segfaultvicta/advent/pkg/seventeen"
	"github.com/segfaultvicta/advent/pkg/sixteen"
)

func main() {
	var day, side, year, infile, instring string
	var lines []string

	flag.StringVar(&day, "day", "1", "index of puzzle we're solving")
	flag.StringVar(&year, "year", "2015", "the year the puzzle belongs to")
	flag.StringVar(&side, "side", "A", "A for side A, B for side B")
	flag.StringVar(&infile, "file", "", "input file to use other than the default")
	flag.StringVar(&instring, "in", "", "input to give to the puzzle, lines seperated by ';'")
	flag.Parse()
	side = strings.ToUpper(strings.Trim(side, " "))
	day = strings.Trim(day, " ")
	year = strings.Trim(year, " ")
	infile = strings.Trim(infile, " ")

	if instring != "" {
		lines = strings.Split(instring, ";")
	} else {
		s := []string{"input/", year, "/", day}

		if infile == "" {
			infile = strings.Join(s, "")
		}

		file, err := os.Open(infile)
		if err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lines = append(lines, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
		defer file.Close()
	}

	fmt.Println(dispatch(year, day, side, lines))
}

func dispatch(year string, day string, side string, lines []string) string {
	switch year {
	case "2015":
		return fifteen.Dispatch(day, side, lines)
	case "2016":
		return sixteen.Dispatch(day, side, lines)
	case "2017":
		return seventeen.Dispatch(day, side, lines)
	case "2018":
		return eighteen.Dispatch(day, side, lines)
	default:
		return "invalid year"
	}
}
