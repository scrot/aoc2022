package day14

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

type loc struct {
	x, y int
}

func (l loc) transform(l2 loc) loc {
	return loc{l.x + l2.x, l.y + l2.y}
}

func (d Day) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	//Draw cave
	sandsrc := loc{500, 0}
	formations := parseInput(buf)
	min, max, grid := buildGrid(formations, sandsrc)

	//Drop sand
	var grains int
	var void bool
	for !void {
		grains++
		grid, void = dropSand(sandsrc, min, max, grid)
	}

	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}

	// Draw cave with floor
	floor := []loc{{min.x - max.y, max.y + 2}, {max.x + max.y, max.y + 2}}
	formations = append(formations, floor)
	min, max, grid = buildGrid(formations, sandsrc)

	var grains2 int
	var hitsrc bool
	for !hitsrc {
		grains2++
		grid, hitsrc = dropSand(sandsrc, min, max, grid)
	}

	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}

	log.Printf("Answer part I: %d\n", grains-1)
	log.Printf("Answer part II: %d\n", grains2)
}

func parseInput(input *bufio.Scanner) [][]loc {
	var formations [][]loc
	for input.Scan() {
		l := input.Text()
		ps := strings.Split(l, " -> ")
		var points []loc
		for _, p := range ps {
			xy := strings.Split(p, ",")
			x, _ := strconv.Atoi(xy[0])
			y, _ := strconv.Atoi(xy[1])
			points = append(points, loc{x, y})
		}
		formations = append(formations, points)
	}

	return formations
}

func buildGrid(formations [][]loc, ss loc) (loc, loc, [][]rune) {
	// Find grid edges
	var (
		min = loc{ss.x, ss.y}
		max = loc{ss.x, ss.y}
	)
	for _, f := range formations {
		for _, p := range f {
			if p.x > max.x {
				max.x = p.x
			}

			if p.x < min.x {
				min.x = p.x
			}

			if p.y > max.y {
				max.y = p.y
			}

			if p.y < min.y {
				min.y = p.y
			}
		}
	}

	// Init an empty grid
	var grid [][]rune
	for i := 0; i <= max.y-min.y; i++ {
		var col []rune
		for j := 0; j <= max.x-min.x; j++ {
			col = append(col, '.')
		}
		grid = append(grid, col)
	}

	// Draw formations
	for _, f := range formations {
		for i := 0; i < len(f)-1; i++ {
			p1, p2 := f[i], f[i+1]
			dx := p2.x - p1.x
			dy := p2.y - p1.y
			for p1 != p2 {
				// draw step
				grid[p1.y-min.y][p1.x-min.x] = '#'
				step := loc{step(dx), step(dy)}
				p1 = p1.transform(step)
			}
			grid[p2.y-min.y][p2.x-min.x] = '#'
		}
	}

	// Add sand source
	grid[ss.y-min.y][ss.x-min.x] = '+'

	return min, max, grid
}

func dropSand(ss, min, max loc, grid [][]rune) ([][]rune, bool) {
	ss = loc{ss.x - min.x, ss.y - min.y}
	grain := ss

	for {
		down := grain.transform(loc{0, 1})
		left := grain.transform(loc{-1, 1})
		right := grain.transform(loc{1, 1})

		//Grain falls in void
		if left.x < 0 || right.x > max.x-min.x {
			return grid, true
		}

		switch {
		case grid[down.y][down.x] != '#' && grid[down.y][down.x] != 'o':
			grain = down
		case grid[left.y][left.x] != '#' && grid[left.y][left.x] != 'o':
			grain = left
		case grid[right.y][right.x] != '#' && grid[right.y][right.x] != 'o':
			grain = right
		default:
			grid[grain.y][grain.x] = 'o'

			//Grain hits sand source
			if grain == ss {
				return grid, true
			}
			return grid, false
		}
	}
}

func step(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}
