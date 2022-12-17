package day10

import (
	"bufio"
	"fmt"
	"log"
	"strings"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

func (d Day) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	signals := newSignals()
	image := newImage()

	x := 1
	var cycle int

	for buf.Scan() {
		l := buf.Text()

		switch l {
		case "noop":
			image = drawPixel(image, cycle, x)
			cycle++
			signals = updateSignals(signals, cycle, x)
		default:
			for i := 0; i < 2; i++ {
				image = drawPixel(image, cycle, x)
				cycle++
				signals = updateSignals(signals, cycle, x)
			}
			x += parseAddr(l)
		}
	}

	log.Printf("Answer part I: %d", sumSignals(signals))
	log.Printf("Answer part II:\n%s", printImage(image))
}

func newImage() [][]rune {
	var image [][]rune
	for i := 0; i < 6; i++ {
		image = append(image, []rune(strings.Repeat(".", 40)))
	}
	return image
}

func printImage(image [][]rune) string {
	b := strings.Builder{}
	for _, row := range image {
		b.WriteString(fmt.Sprintf("%c\n", row))
	}
	return b.String()
}

func newSignals() map[int]int {
	signals := make(map[int]int)
	for i := 20; i <= 220; i += 40 {
		signals[i] = 0
	}
	return signals
}

func updateSignals(signals map[int]int, cycle, x int) map[int]int {
	if _, ok := signals[cycle]; ok {
		signals[cycle] = x
	}
	return signals
}

func sumSignals(signals map[int]int) int {
	var sum int
	for k, v := range signals {
		sum += k * v
	}
	return sum
}

func parseAddr(l string) int {
	var v int
	fmt.Sscanf(l, "addx %d", &v)
	return v
}

func drawPixel(image [][]rune, cycle, x int) [][]rune {
	row, column := int(cycle/40), int(cycle%40)
	if column >= x-1 && column <= x+1 {
		image[row][column] = '#'
	}
	return image
}
