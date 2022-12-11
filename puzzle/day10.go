package puzzle

import (
	"bufio"
	"fmt"
	"log"
)

type Day10 struct {
	*Day
}

func (d Day10) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	x := 1
	var cycle int
	signals := map[int]int{
		20:  0,
		60:  0,
		100: 0,
		140: 0,
		180: 0,
		220: 0,
	}

	for buf.Scan() {
		l := buf.Text()

		switch l {
		case "noop":
			cycle++
			updateSignals(&signals, cycle, x)
		default:
			cycle++
			updateSignals(&signals, cycle, x)
			cycle++
			updateSignals(&signals, cycle, x)
			var v int
			fmt.Sscanf(l, "addx %d", &v)
			x += v
		}

		log.Printf("cycle %d: %d\n", cycle, x)
	}

	var sum int
	for k, v := range signals {
		sum += k * v
	}
	log.Println(signals)

	log.Printf("Answer part I: %d", sum)
	log.Printf("Answer part II: %d", 0)
}

func updateSignals(signals *map[int]int, cycle, x int) {
	if _, ok := (*signals)[cycle]; ok {
		(*signals)[cycle] = x
	}
}
