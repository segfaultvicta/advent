package main

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
	n_conv, n_err := strconv.Atoi(a)
	n_found := false
	i, ok := wires[a]
	if (i >= 0) && ok {
		n = wires[a]
		n_found = true
	} else {
		if n_err == nil {
			n = uint16(n_conv)
			n_found = true
		}
	}

	if n_found == false {
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
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p & q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "OR":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p | q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "LSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p << q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "RSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p >> q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "NOT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				n, n_err := resolve(r["b"], wires)
				if n_err == nil {
					res := ^n
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		default:
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}

				n, err := resolve(r["a"], wires)
				if err == nil {
					wires[r["out"]] = n
					return true
				} else {
					return false
				}
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
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p & q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "OR":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p | q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "LSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p << q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "RSHIFT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				p, p_err := resolve(r["a"], wires)
				q, q_err := resolve(r["b"], wires)
				if p_err == nil && q_err == nil {
					res := p >> q
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		case "NOT":
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}
				n, n_err := resolve(r["b"], wires)
				if n_err == nil {
					res := ^n
					wires[r["out"]] = res
					return true
				} else {
					return false
				}
			}
		default:
			fn = func(wires map[string]uint16) bool {
				_, o_err := resolve(r["out"], wires)
				if o_err == nil {
					return false
				}

				n, err := resolve(r["a"], wires)
				if err == nil {
					wires[r["out"]] = n
					return true
				} else {
					return false
				}
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
