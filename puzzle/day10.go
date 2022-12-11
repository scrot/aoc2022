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
	var image [][]rune
	for i := 0; i < 6; i++ {
		var row []rune
		for i := 0; i < 40; i++ {
			row = append(row, '.')
		}
		image = append(image, row)
	}

	for buf.Scan() {
		l := buf.Text()

		switch l {
		case "noop":
			if cycle%40 >= x-1 && cycle%40 <= x+1 {
				drawPixel(&image, '#', cycle)
			}
			cycle++
			updateSignals(&signals, cycle, x)
		default:
			if cycle%40 >= x-1 && cycle%40 <= x+1 {
				drawPixel(&image, '#', cycle)
			}
			cycle++
			updateSignals(&signals, cycle, x)
			if cycle%40 >= x-1 && cycle%40 <= x+1 {
				drawPixel(&image, '#', cycle)
			}
			cycle++
			updateSignals(&signals, cycle, x)
			var v int
			fmt.Sscanf(l, "addx %d", &v)
			x += v
		}
		fmt.Println(cycle)
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

func drawPixel(image *[][]rune, pixel rune, cycle int) {
	row, column := int(cycle/40), cycle%40
	(*image)[row][column] = pixel

	log.Printf("update row %d col %d\n", row, column)
	for _, row := range *image {
		log.Println(string(row))
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
