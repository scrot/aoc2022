package day15

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

type loc struct {
	x, y int
}

func (l loc) diff(l2 loc) loc {
	diff := loc{l.x - l2.x, l.y - l2.y}
	return diff.abs()
}

func (l loc) abs() loc {
	if l.x < 0 {
		l.x = -l.x
	}
	if l.y < 0 {
		l.y = -l.y
	}
	return l
}

type sensor struct {
	coord  loc
	beacon loc
}

func (s sensor) manhattan() int {
	diff := s.coord.diff(s.beacon)
	return diff.x + diff.y
}

func (d Day) Solve() {
	// Parse input
	index, min, max, sensors := parseInput(d.Dataset)

	// Sensor coverage
	coverage := coverage(sensors)

	// Draw coverage
	grid := draw(min, max, coverage, sensors)

	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}

	var p1 int
	for _, f := range grid[index+10] {
		if f == '#' {
			p1++
		}
	}

	log.Printf("Answer part I: %d", p1)
	log.Printf("Answer part II: %d", 0)
}

func parseInput(input io.ReadCloser) (int, loc, loc, []sensor) {
	buf := bufio.NewScanner(input)
	defer input.Close()

	var (
		min = loc{math.MaxInt, math.MaxInt}
		max = loc{math.MinInt, math.MinInt}
	)

	var index int
	var sensors []sensor
	for buf.Scan() {
		l := buf.Text()

		var c, b loc
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &c.x, &c.y, &b.x, &b.y)
		s := sensor{c, b}

		if s.coord.x-s.manhattan() < min.x {
			min.x = s.coord.x - s.manhattan()
		}

		if s.coord.x+s.manhattan() > max.x {
			max.x = s.coord.x + s.manhattan()
		}

		if s.coord.y-s.manhattan() < min.y {
			index = s.coord.y + s.manhattan()
			min.y = s.coord.y - s.manhattan()
		}

		if s.coord.y+s.manhattan() > max.y {
			max.y = s.coord.y + s.manhattan()
		}

		sensors = append(sensors, s)
	}

	return index, min, max, sensors
}

func coverage(sensors []sensor) []loc {
	var coverage []loc

	for _, sensor := range sensors {

		dist := sensor.manhattan()

		// Coverage top-half
		width := 0
		for i := dist; i >= 0; i-- {
			for j := -width; j <= width; j++ {
				signal := loc{sensor.coord.x + i, sensor.coord.y + j}
				coverage = append(coverage, signal)
			}
			width++
		}

		// Coverage bottom-half (excl. row 0)
		width = 0
		for i := -dist; i < 0; i++ {
			for j := -width; j <= width; j++ {
				signal := loc{sensor.coord.x + i, sensor.coord.y + j}
				coverage = append(coverage, signal)
			}
			width++
		}
	}

	return coverage
}

func draw(min, max loc, coverage []loc, sensors []sensor) [][]rune {
	var grid [][]rune
	for i := min.y; i <= max.y; i++ {
		var row []rune
		for j := min.x; j <= max.x; j++ {
			row = append(row, '.')
		}
		grid = append(grid, row)
	}
	for _, signal := range coverage {
		grid[signal.y+min.abs().y][signal.x+min.abs().x] = '#'
	}

	for _, sensor := range sensors {
		grid[sensor.beacon.y+min.abs().y][sensor.beacon.x+min.abs().x] = 'B'
		grid[sensor.coord.y+min.abs().y][sensor.coord.x+min.abs().x] = 'S'
	}

	return grid
}
