package eighteen

import (
	"fmt"
	"regexp"
	"sort"
	"time"

	u "github.com/segfaultvicta/advent/pkg"
)

type guardAction int

const (
	begins guardAction = 1 + iota
	sleeps
	wakes
)

type guardEvent struct {
	id     int
	action guardAction
	time   time.Time
}

func (e guardEvent) String() string {
	switch e.action {
	case begins:
		return fmt.Sprintf("Guard #%d begins shift at %s", e.id, e.time.Format("2006-01-02 15:04"))
	case sleeps:
		return fmt.Sprintf("Guard #%d goes to sleep at %s", e.id, e.time.Format("2006-01-02 15:04"))
	case wakes:
		return fmt.Sprintf("Guard #%d wakes up at %s", e.id, e.time.Format("2006-01-02 15:04"))
	}
	return "somethin's fucky"
}

func day4sideA(lines []string) string {
	re := regexp.MustCompile(`\[(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2}) (?P<hour>\d{2}):(?P<minute>\d{2})\] ((?P<begins>Guard #(?P<guardId>\d+) begins shift)|(?P<wakes>wakes up)|(?P<sleeps>falls asleep))`)
	events := make([]guardEvent, 0)
	for _, line := range lines {

		params := u.Parse(re, line)
		eventTime := time.Date(
			u.I(params["year"]),
			time.Month(u.I(params["month"])),
			u.I(params["day"]),
			u.I(params["hour"]),
			u.I(params["minute"]),
			0,
			0,
			time.UTC)
		if _, ok := params["begins"]; ok {
			events = append(events, guardEvent{action: begins, id: u.I(params["guardId"]), time: eventTime})
		}
		if _, ok := params["sleeps"]; ok {
			events = append(events, guardEvent{action: sleeps, time: eventTime})
		}
		if _, ok := params["wakes"]; ok {
			events = append(events, guardEvent{action: wakes, time: eventTime})
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time.Before(events[j].time)
	})

	for i, event := range events {
		if event.id == 0 {
			events[i].id = events[i-1].id
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].id == events[j].id {
			return events[i].time.Before(events[j].time)
		}
		return events[i].id < events[j].id
	})

	guardIDs := collectGuardIDs(events)
	sleepiestGuard := 0
	sleepiestGuardMinutes := 0

	for _, id := range guardIDs {
		filteredEvents := filterEventsByID(events, id)
		minutesAsleep := countMinutesAsleep(filteredEvents)
		if minutesAsleep > sleepiestGuardMinutes {
			sleepiestGuard = id
			sleepiestGuardMinutes = minutesAsleep
		}
	}

	sleepiestGuardEvents := filterEventsByID(events, sleepiestGuard)
	sleepiestMinute := 0
	sleepiestMinuteSleeps := 0

	for minute := 0; minute < 60; minute++ {
		sleeps := eventsInWhichAsleep(sleepiestGuardEvents, minute)
		if sleeps > sleepiestMinuteSleeps {
			sleepiestMinute = minute
			sleepiestMinuteSleeps = sleeps
		}
	}

	return fmt.Sprintf("Guard #%d spent the longest asleep, at %d minutes total.\nSleepiest minute: %d, guard was asleep %d times, checksum is %d.\n", sleepiestGuard, sleepiestGuardMinutes, sleepiestMinute, sleepiestMinuteSleeps, sleepiestMinute*sleepiestGuard)
}

func day4sideB(lines []string) string {
	re := regexp.MustCompile(`\[(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2}) (?P<hour>\d{2}):(?P<minute>\d{2})\] ((?P<begins>Guard #(?P<guardId>\d+) begins shift)|(?P<wakes>wakes up)|(?P<sleeps>falls asleep))`)
	events := make([]guardEvent, 0)
	for _, line := range lines {

		params := u.Parse(re, line)
		eventTime := time.Date(
			u.I(params["year"]),
			time.Month(u.I(params["month"])),
			u.I(params["day"]),
			u.I(params["hour"]),
			u.I(params["minute"]),
			0,
			0,
			time.UTC)
		if _, ok := params["begins"]; ok {
			events = append(events, guardEvent{action: begins, id: u.I(params["guardId"]), time: eventTime})
		}
		if _, ok := params["sleeps"]; ok {
			events = append(events, guardEvent{action: sleeps, time: eventTime})
		}
		if _, ok := params["wakes"]; ok {
			events = append(events, guardEvent{action: wakes, time: eventTime})
		}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].time.Before(events[j].time)
	})

	for i, event := range events {
		if event.id == 0 {
			events[i].id = events[i-1].id
		}
	}

	sort.Slice(events, func(i, j int) bool {
		if events[i].id == events[j].id {
			return events[i].time.Before(events[j].time)
		}
		return events[i].id < events[j].id
	})

	guardIDs := collectGuardIDs(events)

	for _, id := range guardIDs {
		sleepiestMinute := 0
		sleepiestMinuteSleeps := 0
		filteredEvents := filterEventsByID(events, id)
		// Which minute did this guard spend asleep the most times?
		for minute := 0; minute < 60; minute++ {
			sleeps := eventsInWhichAsleep(filteredEvents, minute)
			if sleeps > sleepiestMinuteSleeps {
				sleepiestMinute = minute
				sleepiestMinuteSleeps = sleeps
			}
		}
		fmt.Printf("Guard %d spent %d asleep the most - %d times.\n", id, sleepiestMinute, sleepiestMinuteSleeps)
	}

	return "Look at the above output and figure it out, scrub."
}

func eventsInWhichAsleep(events []guardEvent, minute int) int {
	eventsAsleep := 0

	currentlySleeping := false
	fellAsleepAt := 0
	for _, event := range events {
		if currentlySleeping {
			// there'll never be two sleeping actions in a row so this is fine
			if minute >= fellAsleepAt && minute < event.time.Minute() {
				eventsAsleep++
			}

			currentlySleeping = false
		} else {
			if event.action == sleeps {
				currentlySleeping = true
				fellAsleepAt = event.time.Minute()
			}
		}
	}

	return eventsAsleep
}

func countMinutesAsleep(events []guardEvent) int {
	minutesAsleep := 0
	currentlySleeping := false
	fellAsleepAt := 0
	for _, event := range events {
		if currentlySleeping {
			// there'll never be two sleeping actions in a row so this is fine
			minutesAsleep += event.time.Minute() - fellAsleepAt
			fellAsleepAt = 0
			currentlySleeping = false
		} else {
			if event.action == sleeps {
				currentlySleeping = true
				fellAsleepAt = event.time.Minute()
			}
		}
	}
	return minutesAsleep
}

func filterEventsByID(events []guardEvent, id int) []guardEvent {
	filteredEvents := make([]guardEvent, 0)
	for _, event := range events {
		if event.id == id {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents
}

func collectGuardIDs(events []guardEvent) []int {
	allIds := make([]int, 0)
	ids := make([]int, 0)
	for _, event := range events {
		allIds = append(allIds, event.id)
	}
	for _, id := range allIds {
		if !u.CheckIntInSlice(id, ids) {
			ids = append(ids, id)
		}
	}
	return ids
}
