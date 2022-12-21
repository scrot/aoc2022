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

type area struct {
	center                 loc
	xmin, xmax, ymin, ymax loc
}

type sensor struct {
	coord  loc
	beacon loc
}

func (s sensor) contains(l loc) bool {
	a := s.manhattan()
	b := s.manhattan()
	U := 2 * s.manhattan()
	V := 2 * s.manhattan()
	W := loc{l.x - s.coord.x, l.y - s.coord.y}
	abs := loc{W.x * U, W.y * V}.abs()
	if abs.x/a+abs.y/b <= 1 {
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
	sensors := parseInput(d.Dataset)

	// Sensor coverage
	// coverage := coverage(sensors)

	//
	sensors[0].contains(loc{4, 20})
	fmt.Println("Point 4,20 is contained by %v\n", sensors[0])

	log.Printf("Answer part I: %d", 0)
	log.Printf("Answer part II: %d", 0)
}

func parseInput(input io.ReadCloser) []sensor {
	buf := bufio.NewScanner(input)
	defer input.Close()

	var sensors []sensor
	for buf.Scan() {
		l := buf.Text()

		var c, b loc
		fmt.Sscanf(l, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &c.x, &c.y, &b.x, &b.y)
		s := sensor{c, b}

		sensors = append(sensors, s)
	}

	return sensors
}

func coverage(sensors []sensor) []area {
	var coverage []area

	for _, sensor := range sensors {
		dist := sensor.manhattan()
		a := area{
			loc{sensor.coord.x, sensor.coord.y},
			loc{sensor.coord.x - dist, sensor.coord.y},
			loc{sensor.coord.x + dist, sensor.coord.y},
			loc{sensor.coord.x, sensor.coord.y - dist},
			loc{sensor.coord.x, sensor.coord.y + dist},
		}

		coverage = append(coverage, a)
	}

	return coverage
}
