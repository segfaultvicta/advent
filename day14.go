package main

import "fmt"

type Reindeer struct {
	Name             string
	KMPS             int
	Stamina          int
	Nappin           bool
	NapDuration      int
	CurrentNapLength int
	CurrentRunLength int
}

func (r *Reindeer) String() string {
	if r.Nappin == true {
		return fmt.Sprintln(r.Name, "can run at", r.KMPS, "km/s for", r.Stamina, "seconds, is currently napping and on nap tick", r.CurrentNapLength, "of", r.NapDuration)
	} else {
		return fmt.Sprintln(r.Name, "can run at", r.KMPS, "km/s for", r.Stamina, "seconds, is not currently napping, but will need to nap for", r.NapDuration)
	}
}

func (r *Reindeer) Tick() int {
	if r.Nappin {
		r.CurrentNapLength++
		if r.CurrentNapLength == r.NapDuration {
			r.Nappin = false
			r.CurrentNapLength = 0
		}
		return 0
	} else {
		r.CurrentRunLength++
		if r.CurrentRunLength == r.Stamina {
			r.Nappin = true
			r.CurrentRunLength = 0
		}
		return r.KMPS
	}
}

func day14sideA(lines []string) string {
	var reindeerList map[string]*Reindeer
	reindeerList = make(map[string]*Reindeer)
	var name string
	var speed, dur, rest int
	for _, line := range lines {
		// today I learned how sscanf worked! :D
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &dur, &rest)
		reindeerList[name] = &Reindeer{Name: name, KMPS: speed, Stamina: dur, NapDuration: rest, Nappin: false}
	}

	var raceProgress map[*Reindeer]int
	raceProgress = make(map[*Reindeer]int)
	for tick := 0; tick < 2503; tick++ {
		for _, reindeer := range reindeerList {
			raceProgress[reindeer] += reindeer.Tick()
		}
	}

	fmt.Println(raceProgress)

	return "^^^"
}

func day14sideB(lines []string) string {
	var reindeerList map[string]*Reindeer
	reindeerList = make(map[string]*Reindeer)
	var name string
	var speed, dur, rest int
	for _, line := range lines {
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.", &name, &speed, &dur, &rest)
		reindeerList[name] = &Reindeer{Name: name, KMPS: speed, Stamina: dur, NapDuration: rest, Nappin: false}
	}

	var raceProgress map[*Reindeer]int
	raceProgress = make(map[*Reindeer]int)

	var scoreList map[*Reindeer]int
	scoreList = make(map[*Reindeer]int)

	var firstPlace []*Reindeer

	for tick := 0; tick < 2503; tick++ {
		for _, reindeer := range reindeerList {
			raceProgress[reindeer] += reindeer.Tick()
		}
		// find currently winning reindeer
		best := 0
		firstPlace = []*Reindeer{}
		for _, progress := range raceProgress {
			if best < progress {
				best = progress
			}
		}
		for reindeer, progress := range raceProgress {
			if progress == best {
				firstPlace = append(firstPlace, reindeer)
			}
		}

		for _, reindeer := range firstPlace {
			scoreList[reindeer]++
		}
	}

	fmt.Println(scoreList)

	return "^^^"
}
