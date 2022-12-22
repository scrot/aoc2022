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

func (s sensor) contains(l loc) bool {
	// distance point from center
	dist := loc{l.x - s.coord.x, l.y - s.coord.y}.abs()

	// dist point smaller than edge distance then in diamond
	if dist.x+dist.y < s.manhattan() {
		return true
	}
	return false
}

func (s sensor) manhattan() int {
	diff := s.coord.diff(s.beacon)
	return diff.x + diff.y
}

func (d Day) Solve() {
	// Parse input
	_, _, sensors := parseInput(d.Dataset)

	p1 := 2000000
	m := make(map[loc]bool)
	for _, s := range sensors {
		y := loc{s.coord.x, p1}
		diff := y.diff(s.coord)

		// Add points overlap
		if width := s.manhattan() - diff.y; width >= 0 {
			// fmt.Printf("Sensor: %v, Dist: %d, Width: %d\n", s.coord, s.manhattan(), width)
			for x := -width; x <= width; x++ {
				l := loc{s.coord.x + x, p1}
				m[l] = true
			}
		}
	}

	for _, s := range sensors {
		if _, ok := m[s.beacon]; ok {
			delete(m, s.beacon)
		}
	}

	log.Printf("Answer part I: %d", len(m))
	log.Printf("Answer part II: %d", 0)
}

func parseInput(input io.ReadCloser) (loc, loc, []sensor) {
	buf := bufio.NewScanner(input)
	defer input.Close()

	var (
		min = loc{math.MaxInt, math.MaxInt}
		max = loc{math.MaxInt, math.MaxInt}
	)

	var sensors []sensor
	for buf.Scan() {
		l := buf.Text()

		var c, b loc
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &c.x, &c.y, &b.x, &b.y)
		s := sensor{c, b}

		if c.x < min.x {
			min.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
		if c.x > max.x {
			max.x = c.y
		}
		if c.y > max.y {
			max.y = c.y
		}

		sensors = append(sensors, s)
	}

	return min, max, sensors
}
