package puzzle

import (
	"bufio"
	"fmt"
	"log"
	"strings"
)

type Day10 struct {
	*Day
}

func (d Day10) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	x := 1
	var cycle int

	signals := make(map[int]int)
	for i := 20; i <= 220; i += 40 {
		signals[i] = 0
	}

	var image [][]rune
	for i := 0; i < 6; i++ {
		image = append(image, []rune(strings.Repeat(".", 40)))
	}

	for buf.Scan() {
		l := buf.Text()

		switch l {
		case "noop":
			image = drawPixel(image, cycle, x)
			cycle++
			updateSignals(&signals, cycle, x)
		default:
			for i := 0; i < 2; i++ {
				image = drawPixel(image, cycle, x)
				cycle++
				updateSignals(&signals, cycle, x)
			}
			var v int
			fmt.Sscanf(l, "addx %d", &v)
			x += v
		}
	}

	var sum int
	for k, v := range signals {
		sum += k * v
	}

	log.Printf("Answer part I: %d", sum)
	log.Println("Answer part II:")
	for _, row := range image {
		log.Println(string(row))
	}
}

func updateSignals(signals *map[int]int, cycle, x int) {
	if _, ok := (*signals)[cycle]; ok {
		(*signals)[cycle] = x
	}
}

func drawPixel(image [][]rune, cycle, x int) [][]rune {
	row, column := int(cycle/40), int(cycle%40)
	if column >= x-1 && column <= x+1 {
		image[row][column] = '#'
	}
	return image
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
