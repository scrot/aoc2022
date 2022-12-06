package puzzle

import (
	"bufio"
	"fmt"
	"log"
)

type Day5 struct {
	*Day
}

type stacks [9][]rune

func (d Day5) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var layout []string

	for buf.Scan() {
		l := buf.Text()
		if len(l) == 0 {
			break
		}
		layout = append(layout, l)
	}

	stacks9000 := parseStackLayout(layout)
	stacks9001 := parseStackLayout(layout)

	for buf.Scan() {
		l := buf.Text()

		var m, f, t int
		fmt.Sscanf(l, "move %d from %d to %d\n", &m, &f, &t)
		// log.Printf("move %d, from %d, to %d", inst[0], inst[1], inst[2])
		stacks9000 = moveCrates(stacks9000, m, f, t, false)
		stacks9001 = moveCrates(stacks9001, m, f, t, true)
	}

	top9000 := topCrates(stacks9000)
	top9001 := topCrates(stacks9001)

	log.Printf("Answer part I: %s\n", string(top9000))
	log.Printf("Answer part II: %s\n", string(top9001))
}

func parseStackLayout(layout []string) stacks {
	var s stacks
	for i := len(layout) - 2; i >= 0; i-- {
		var p int
		for j := 0; j < len(layout[i]); j += 4 {
			symbol := rune(layout[i][j+1])
			if symbol != ' ' {
				s[p] = append(s[p], symbol)
			}
			p++
		}
		// log.Printf("%s\n", layout[i])
	}

	return s
}

func moveCrates(stacks stacks, amount, from, to int, is9001 bool) stacks {
	pop := len(stacks[from-1]) - amount
	crates := stacks[from-1][pop:len(stacks[from-1])]
	// log.Printf("crates to move: %c", crates)

	stacks[from-1] = stacks[from-1][:pop]

	if is9001 {
		for _, crate := range crates {
			stacks[to-1] = append(stacks[to-1], crate)
		}
	} else {
		for i := len(crates) - 1; i >= 0; i-- {
			stacks[to-1] = append(stacks[to-1], crates[i])
		}
	}

	return stacks
}

func topCrates(stacks stacks) []rune {
	var top []rune
	for _, stack := range stacks {
		top = append(top, stack[len(stack)-1])
	}
	return top
}
