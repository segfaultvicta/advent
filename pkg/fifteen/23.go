package fifteen

import (
	"fmt"
	"strconv"
	"strings"
)

func day23sideA(lines []string) string {
	instructions := [][]string{}
	for i, line := range lines {
		split := strings.Split(line, " ")
		fmt.Printf("%d: %v\n", i, split)
		instructions = append(instructions, split)
	}

	mc := 0
	registers := map[string]int{"a": 1, "b": 0}

	//reader := bufio.NewReader(os.Stdin)

	for mc <= 45 {
		instruction := instructions[mc]
		op := instruction[0]
		//fmt.Printf("[%d] %v a(%d) b(%d)\n", mc, instruction, registers["a"], registers["b"])
		//reader.ReadString('\n')
		switch {
		case op == "jmp":
			// jmp offset is a jump; it continues with the instruction offset away relative to itself
			offset, _ := strconv.Atoi(instruction[1])
			mc += offset
		case op == "jie":
			// jie r, offset is like jmp, but only jumps if register r is even ("jump if even").
			register := instruction[1][0:1]
			offset, _ := strconv.Atoi(instruction[2])
			//fmt.Printf("Jumping to %d if %d from %v is even...\n", mc+offset, registers[register], register)
			if registers[register]%2 == 0 {
				mc += offset
			} else {
				mc++
			}
		case op == "jio":
			// jio r, offset is like jmp, but only jumps if register r is 1 ("jump if one", not odd).
			register := instruction[1][0:1]
			offset, _ := strconv.Atoi(instruction[2])
			if registers[register] == 1 {
				mc += offset
			} else {
				mc++
			}
		case op == "inc":
			// inc r increments register r, adding 1 to it, then continues with the next instruction.
			register := instruction[1]
			registers[register]++
			mc++
		case op == "tpl":
			// tpl r sets register r to triple its current value, then continues with the next instruction.
			register := instruction[1]
			registers[register] *= 3
			mc++
		case op == "hlf":
			// hlf r sets register r to half its current value, then continues with the next instruction.
			register := instruction[1]
			registers[register] /= 2
			mc++
		}
	}

	return fmt.Sprintf("Register a: %d Register b: %d", registers["a"], registers["b"])
}

func day23sideB(lines []string) string {
	return "n/i"
}
