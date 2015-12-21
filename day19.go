package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

type Parser struct {
	Rules  map[string][]string
	Target string
}

type ParseResult struct {
	Index int
	Token string
	To    []string
}

func (p *ParseResult) String() string {
	return "<{" + strconv.Itoa(p.Index) + "}: " + p.Token + " parses to: " + strings.Join(p.To, ",") + ">"
}

func NewParser(rules map[string][]string, target string) *Parser {
	p := new(Parser)
	p.Rules = rules
	p.Target = target
	return p
}

func MakeParseResult(index int, token string, results []string) *ParseResult {
	p := new(ParseResult)
	p.Index = index
	p.Token = token
	p.To = results
	return p
}

func (p *Parser) BestChunkedDistance(candidate string) int {
	chunks := len(p.Target) / len(candidate) * 2
	splitat := len(candidate) / 2
	best := 9999999
	for i := 0; i < chunks-1; i++ {
		chunk := p.Target[i*splitat : i*splitat+len(candidate)]
		ld := levenshtein.DistanceForStrings([]rune(chunk), []rune(candidate), levenshtein.DefaultOptions)
		if ld < best {
			best = ld
		}
	}
	return best
}

func (p *Parser) Parse(input string) <-chan *ParseResult {
	ch := make(chan *ParseResult)
	length := len(input)
	index := 0
	go func() {
		for index < length {
			// try to chew two runes off the input string, see if that matches anything in precursor
			if index < length-1 {
				chew := input[index : index+2]
				if val, ok := p.Rules[chew]; ok {
					ch <- MakeParseResult(index, chew, val)
					index += 2
				} else {
					chew = input[index : index+1]
					if val, ok := p.Rules[chew]; ok {
						ch <- MakeParseResult(index, chew, val)
						index++
					} else {
						panic("parse error: invalid token " + chew + " at index " + strconv.Itoa(index))
					}
				}
			} else {
				// we're on the last atom!
				chew := input[index : index+1]
				if val, ok := p.Rules[chew]; ok {
					ch <- MakeParseResult(index, chew, val)
					index++
				} else {
					panic("parse error: invalid token " + chew + " at index " + strconv.Itoa(index))
				}
			}
		}
		close(ch)
	}()
	return ch
}

func ReplaceAt(index int, in string, tokenlength int, r string) string {
	p1 := in[0:index]
	p2 := in[index+tokenlength : len(in)]
	return strings.Join([]string{p1, r, p2}, "")
}

func day19sideA(lines []string) string {
	molecule := lines[len(lines)-1]

	lines = lines[0 : len(lines)-2]

	var rules map[string][]string
	rules = make(map[string][]string)

	for _, rx := range lines {
		split := strings.Split(rx, " ")
		//precursors = append(precursors, split[0])
		//results = append(results, split[2])
		if val, ok := rules[split[0]]; ok {
			rules[split[0]] = append(val, split[2])
		} else {
			results := []string{split[2]}
			rules[split[0]] = results
		}
	}

	parser := NewParser(rules, molecule)

	var parsed []*ParseResult
	for parseResult := range parser.Parse(molecule) {
		parsed = append(parsed, parseResult)
	}

	molecules := map[string]int{}

	for _, r := range parsed {
		// replace the token at index of molecule with each possible parse result
		//fmt.Println(r)

		for i := 0; i < len(r.To); i++ {
			//molecules = append(molecules, ReplaceAt(r.Index, molecule, len(r.Token), r.To[i]))
			candidate := ReplaceAt(r.Index, molecule, len(r.Token), r.To[i])
			//fmt.Println(candidate)
			if _, seen := molecules[candidate]; !seen {
				molecules[candidate] = 1
			} else {
				molecules[candidate]++
			}
		}
	}

	for k, v := range molecules {
		if v > 1 {
			fmt.Println(k, v)
		}
	}

	return strconv.Itoa(len(molecules) - 1)
}

func Freeze(in []*ParseResult) (out string) {
	// takes the token of every parse result and emits a string
	for _, r := range in {
		out += r.Token
	}
	return out
}

func (p *Parser) VargHammer(in []*ParseResult, depth, depthLimit int) ([]*ParseResult, int) {
	// unless depthlimit has been exceeded or no valid parse results exist,
	if depth == depthLimit {
		//fmt.Println("depth limit reached, oh no")
		return in, depth
	}
	//fmt.Println("===============================")
	//fmt.Println("SMASH WIF HAMAR!", depth, depthLimit)
	var valid []*ParseResult
	for _, r := range in {
		//fmt.Println("checking validity", r)
		if len(r.To) > 1 {
			valid = append(valid, r)
		}
	}
	if len(valid) == 0 {
		//fmt.Println("base case")
		return in, depth
	}
	if len(valid) < 0 {
		panic("somehow no valid values exist at all wtf")
	}
	// randomly select a token from in, and a valid parse target from that result

	choice := rand.Intn(len(valid))
	choose := valid[choice]
	target := choose.To[rand.Intn(len(choose.To))]
	//fmt.Println("choosing token", choose.Token, "and result", target)
	results := []*ParseResult{}
	// generate a new parse result from parsing that parse target
	for parseResult := range p.Parse(target) {
		//fmt.Println("that result parses to", parseResult)
		results = append(results, parseResult) // think I need range for this because lolreasons
	}
	//in[choice] = nil // avoid horrible memory leak, lol :(
	// replace results into in at position choice
	for j := 0; j < len(results)-1; j++ {
		in = append(in, new(ParseResult)) // append zero-value len(r)-1 times
	}
	copy(in[choice+len(results)-1:], in[choice:]) // shift shit over len(r)-1 times
	for j, x := range results {
		in[choice+j] = x
	}
	//fmt.Println(in)
	foo, bar := p.VargHammer(in, depth+1, depthLimit)
	//fmt.Println("RETURNING: ", foo, bar)
	return foo, bar
}

func day19sideB(lines []string) string {
	// FUCK ALL OF THIS

	//count := 0

	// molecule := lines[len(lines)-1]

	// lines = lines[0 : len(lines)-2]

	// var rules map[string][]string
	// rules = make(map[string][]string)

	// for _, rx := range lines {
	// 	split := strings.Split(rx, " ")
	// 	if val, ok := rules[split[0]]; ok {
	// 		rules[split[0]] = append(val, split[2])
	// 	} else {
	// 		results := []string{split[2]}
	// 		rules[split[0]] = results
	// 	}
	// }

	// parser := NewParser(rules, molecule)

	// begin := "e"

	// var test []*ParseResult
	// for parseResult := range parser.Parse(begin) {
	// 	test = append(test, parseResult)
	// }

	// threshold := 20
	// depthLimit := 80
	// rand.Seed(time.Now().UnixNano())
	// for {
	// 	testParse, depth := parser.VargHammer(test, 0, depthLimit)
	// 	frozen := Freeze(testParse)
	// 	//chunkedbest := parser.BestChunkedDistance(frozen)
	// 	//fmt.Println(chunkedbest)
	// 	if strings.Contains(molecule, frozen) {
	// 		fmt.Println(len(molecule), len(frozen), frozen)
	// 		panic("at the disco")
	// 		if len(frozen) > threshold {
	// 			fmt.Println(testParse, "at depth", depth)
	// 		}
	// 	}
	// }

	molecule := lines[len(lines)-1]
	molecule = strings.Replace(molecule, "Rn", " { ", -1)
	molecule = strings.Replace(molecule, "Ar", " } ", -1)
	molecule = strings.Replace(molecule, "Y", " . ", -1)

	return molecule

}
