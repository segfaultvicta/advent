package fifteen

import (
	"errors"
	"regexp"
	"sort"
	"strconv"
)

var re = regexp.MustCompile("(?P<a>[0-9a-z]*)? ?(?P<op>AND|OR|NOT|LSHIFT|RSHIFT)? ?(?P<b>[0-9a-z]*) -> (?P<out>[a-z]*)")

type circuitFunc func(map[string]uint16) bool

func resolve(a string, wires map[string]uint16) (n uint16, err error) {
	if a == "b" {
		return 16076, nil
	}
	nConv, nErr := strconv.Atoi(a)
	nFound := false
	i, ok := wires[a]
	if (i >= 0) && ok {
		n = wires[a]
		nFound = true
	} else {
		if nErr == nil {
			n = uint16(nConv)
			nFound = true
		}
	}

	if nFound == false {
		err = errors.New(a + " does not resolve")
	} else {
		err = nil
	}
	return n, err
}

func day7sideA(lines []string) string {

	var graph []circuitFunc

	var wires map[string]uint16
	wires = make(map[string]uint16)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		r := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 {
				r[name] = match[i]
			}
		}

		var fn circuitFunc
		switch r["op"] {
		case "AND":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p & q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "OR":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p | q
					wires[r["out"]] = res
					return true
				}
				return false

			}
		case "LSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p << q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "RSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p >> q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "NOT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				n, nErr := resolve(r["b"], wires)
				if nErr == nil {
					res := ^n
					wires[r["out"]] = res
					return true
				}
				return false
			}
		default:
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}

				n, err := resolve(r["a"], wires)
				if err == nil {
					wires[r["out"]] = n
					return true
				}
				return false
			}
		}
		graph = append(graph, fn)
	}

	running := true

	for running {
		finished := true
		for _, gate := range graph {
			if gate(wires) {
				finished = false
			}
		}
		if finished == true {
			running = false
		}
	}

	var keys []string
	for k := range wires {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ret := ""
	for _, k := range keys {
		ret += k + ": " + strconv.Itoa(int(wires[k])) + "\n"
	}
	return ret
}

func day7sideB(lines []string) string {
	var graph []circuitFunc

	var wires map[string]uint16
	wires = make(map[string]uint16)

	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		r := make(map[string]string)
		for i, name := range re.SubexpNames() {
			if i != 0 {
				r[name] = match[i]
			}
		}

		var fn circuitFunc
		switch r["op"] {
		case "AND":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p & q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "OR":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p | q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "LSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p << q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "RSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				p, pErr := resolve(r["a"], wires)
				q, qErr := resolve(r["b"], wires)
				if pErr == nil && qErr == nil {
					res := p >> q
					wires[r["out"]] = res
					return true
				}
				return false
			}
		case "NOT":
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}
				n, nErr := resolve(r["b"], wires)
				if nErr == nil {
					res := ^n
					wires[r["out"]] = res
					return true
				}
				return false
			}
		default:
			fn = func(wires map[string]uint16) bool {
				_, oErr := resolve(r["out"], wires)
				if oErr == nil {
					return false
				}

				n, err := resolve(r["a"], wires)
				if err == nil {
					wires[r["out"]] = n
					return true
				}
				return false
			}
		}
		graph = append(graph, fn)
	}

	running := true

	for running {
		finished := true
		for _, gate := range graph {
			if gate(wires) {
				finished = false
			}
		}
		if finished == true {
			running = false
		}
	}

	var keys []string
	for k := range wires {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	ret := ""
	for _, k := range keys {
		ret += k + ": " + strconv.Itoa(int(wires[k])) + "\n"
	}
	return ret
}
