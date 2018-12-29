package eighteen

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	u "github.com/segfaultvicta/advent/pkg"
)

func day7sideA(lines []string) string {
	re := regexp.MustCompile(`Step (?P<prerequisite>.) must be finished before step (?P<target>.) can begin.`)
	blockedBy := make(map[string][]string, 0)
	for _, line := range lines {
		params := u.Parse(re, line)
		blockedBy[params["target"]] = append(blockedBy[params["target"]], params["prerequisite"])
	}
	blockedBy["E"] = []string{}
	blockedBy["I"] = []string{}
	blockedBy["V"] = []string{}
	blockedBy["Z"] = []string{}

	steps := make([]string, 0)
	lastStep := ""
	for len(steps) <= 25 {
		steps, blockedBy = path(steps, blockedBy)
		lastStep = steps[len(steps)-1]
		delete(blockedBy, lastStep)
	}

	return strings.Join(steps, "")
}

func path(acc []string, blockedBy map[string][]string) ([]string, map[string][]string) {
	//check to see if anything needs to be added to unblocked
	unblocked := unblocked(acc, blockedBy)

	sort.Strings(unblocked)

	return append(acc, unblocked[0]), blockedBy
}

func unblocked(path []string, blockedBy map[string][]string) []string {
	unblocked := make([]string, 0)
	for blocked, bySlice := range blockedBy {
		// remove everything from 'by' that is in 'acc'
		blockedNum := len(bySlice)
		if blockedNum > 0 {
			for _, by := range bySlice {
				if u.CheckStringInSlice(by, path) {
					blockedNum--
				}
			}
			if blockedNum == 0 {
				unblocked = append(unblocked, blocked)
			}
		} else {
			unblocked = append(unblocked, blocked)
		}
	}
	sort.Strings(unblocked)
	return unblocked
}

type taskManager struct {
	tick      int
	workers   []*worker
	path      []string
	blockedBy map[string][]string
}

func taskDuration(task string) int {
	if task == " " || task == "" {
		return 1
	}
	return int(rune(task[0]) - 4)
	//return int(rune(task[0]) - 64)
}

func (m *taskManager) taskTick() {
	//fmt.Printf("==========TICK %d==========\n", m.tick)
	for _, worker := range m.workers {
		if worker.timer == 0 {
			// worker is done with the task! add task to path
			if worker.task != " " && worker.task != "" {
				//fmt.Printf("Worker %d finished %s!\n", i, worker.task)
				m.path = append(m.path, worker.task)
				delete(m.blockedBy, worker.task)
				worker.task = " "
			}
			// worker is not currently on a task, try to assign it
			//fmt.Printf("Trying to assign a task to Worker %d...\n", i)
			unblocked := unblocked(m.path, m.blockedBy)
			//fmt.Printf("Current unblocked tasks: %v\n", unblocked)
			if len(unblocked) > 0 {
				for _, candidateTask := range unblocked {
					for _, taskCheck := range m.workers {
						if taskCheck.task == candidateTask {
							candidateTask = " "
						}
					}
					//fmt.Printf("Found work to do - %s\n", candidateTask)
					worker.task = candidateTask
					worker.timer = taskDuration(candidateTask)
					if candidateTask != " " {
						break
					}
				}
			} else {
				worker.timer = 1
			}
		}
		worker.timer--
	}
	//time.Sleep(1000 * time.Millisecond)
	fmt.Printf("%d:\t%s (%d)\t%s (%d)\t%s (%d)\t%s (%d)\t%s (%d)\t%s\t\t\topen: %s\n", m.tick, m.workers[0].task, m.workers[0].timer, m.workers[1].task, m.workers[1].timer, m.workers[2].task, m.workers[2].timer, m.workers[3].task, m.workers[3].timer, m.workers[4].task, m.workers[4].timer, strings.Join(m.path, ""), strings.Join(unblocked(m.path, m.blockedBy), ""))
	m.tick++
}

type worker struct {
	timer int
	task  string
}

func day7sideB(lines []string) string {
	re := regexp.MustCompile(`Step (?P<prerequisite>.) must be finished before step (?P<target>.) can begin.`)
	blockedBy := make(map[string][]string, 0)
	for _, line := range lines {
		params := u.Parse(re, line)
		blockedBy[params["target"]] = append(blockedBy[params["target"]], params["prerequisite"])
	}
	blockedBy["E"] = []string{}
	blockedBy["I"] = []string{}
	blockedBy["V"] = []string{}
	blockedBy["Z"] = []string{}
	//blockedBy["C"] = []string{}

	steps := []string{}
	workers := []*worker{&worker{0, ""}, &worker{0, ""}, &worker{0, ""}, &worker{0, ""}, &worker{0, ""}}
	//workers := []*worker{&worker{0, ""}, &worker{0, ""}}
	taskManager := taskManager{0, workers, []string{}, blockedBy}
	for len(steps) < 26 {
		//fmt.Println(steps)
		taskManager.taskTick()
		steps = taskManager.path
		//fmt.Println(steps)
	}

	// 908 is too high

	return fmt.Sprintf("took %d ticks to complete\n", taskManager.tick-1)
}
