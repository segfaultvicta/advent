package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/fighterlyt/permutation"
)

func NewPerson(name string) *Person {
	var opinions map[string]int
	opinions = make(map[string]int)
	return &Person{name, opinions}
}

type Person struct {
	Name     string
	Opinions map[string]int
}

func (p *Person) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("\n" + p.Name + " has the following opinions:\n")
	for name, opinion := range p.Opinions {
		buffer.WriteString("\t" + name + "\t\t" + strconv.Itoa(opinion) + "\n")
	}
	return buffer.String()
}

func (p *Person) setOpinionAbout(name string, to int) {
	p.Opinions[name] = to
}

func selectPerson(person string, list []*Person) *Person {
	for _, e := range list {
		if e.Name == person {
			return e
		}
	}
	return nil
}

func calculateHappiness(arrangement []*Person) (happiness int) {
	for i, p := range arrangement {
		switch {
		case i == 0:
			happiness += p.Opinions[arrangement[len(arrangement)-1].Name] + p.Opinions[arrangement[1].Name]
		case i == (len(arrangement) - 1):
			happiness += p.Opinions[arrangement[0].Name] + p.Opinions[arrangement[i-1].Name]
		default:
			happiness += p.Opinions[arrangement[i-1].Name] + p.Opinions[arrangement[i+1].Name]
		}
	}
	return happiness
}

func day13sideA(lines []string) string {
	re := regexp.MustCompile("([A-Za-z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([A-Za-z]+).")
	var guestList = []*Person{NewPerson("Alice"), NewPerson("Bob"), NewPerson("Carol"), NewPerson("David"), NewPerson("Eric"), NewPerson("Frank"), NewPerson("George"), NewPerson("Mallory")}
	for _, line := range lines {
		pieces := re.FindStringSubmatch(line)
		a := pieces[1]
		num, _ := strconv.Atoi(pieces[3])
		b := pieces[4]
		if pieces[2] == "lose" {
			num *= -1
		}
		personA := selectPerson(a, guestList)
		personA.setOpinionAbout(b, num)
	}

	less := func(i, j interface{}) bool {
		p := i.(*Person)
		q := j.(*Person)
		return p.Name < q.Name
	}
	permuter, _ := permutation.NewPerm(guestList, less)
	var permutations [][]*Person

	for i, err := permuter.Next(); err == nil; i, err = permuter.Next() {
		//fmt.Printf("%3d permutation: %v, %d left\n", permuter.Index()-1, i.([]*Person), permuter.Left())
		permutations = append(permutations, i.([]*Person))
	}

	best := 0
	for _, p := range permutations {
		if result := calculateHappiness(p); result > best {
			best = result
		}
	}
	return strconv.Itoa(best)
}

func day13sideB(lines []string) string {
	re := regexp.MustCompile("([A-Za-z]+) would (gain|lose) ([0-9]+) happiness units by sitting next to ([A-Za-z]+).")
	var guestList = []*Person{NewPerson("Alice"), NewPerson("Bob"), NewPerson("Carol"), NewPerson("David"), NewPerson("Eric"), NewPerson("Frank"), NewPerson("George"), NewPerson("Mallory"), NewPerson("Sadsack")}
	for _, line := range lines {
		pieces := re.FindStringSubmatch(line)
		a := pieces[1]
		num, _ := strconv.Atoi(pieces[3])
		b := pieces[4]
		if pieces[2] == "lose" {
			num *= -1
		}
		personA := selectPerson(a, guestList)
		personA.setOpinionAbout(b, num)
	}

	for _, guest := range guestList {
		guest.setOpinionAbout("Sadsack", 0)
	}

	fmt.Println(guestList)

	less := func(i, j interface{}) bool {
		p := i.(*Person)
		q := j.(*Person)
		return p.Name < q.Name
	}
	permuter, _ := permutation.NewPerm(guestList, less)
	var permutations [][]*Person

	for i, err := permuter.Next(); err == nil; i, err = permuter.Next() {
		//fmt.Printf("%3d permutation: %v, %d left\n", permuter.Index()-1, i.([]*Person), permuter.Left())
		permutations = append(permutations, i.([]*Person))
	}

	best := 0
	for _, p := range permutations {
		if result := calculateHappiness(p); result > best {
			best = result
		}
	}
	return strconv.Itoa(best)
}
