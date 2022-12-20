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
	//Parse input
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	//Draw cave
	sandsrc := loc{500, 0}
	formations := parseInput(buf)
	dim, grid := buildGrid(formations, sandsrc, 1)

	//Drop sand
	var grains int
	var void bool
	for !void {
		grains++
		grid, void = dropSand(sandsrc, dim, grid)
	}

	for _, row := range grid {
		fmt.Printf("%c\n", row)
	}

	// Draw cave with floor
	width := (dim[3] - dim[2])
	floor := []loc{{dim[0] - width, dim[3] + 2}, {dim[1] + width, dim[3] + 2}}
	formations = append(formations, floor)
	dim2, grid2 := buildGrid(formations, sandsrc, 0)

	for _, row := range grid2 {
		fmt.Printf("%c\n", row)
	}

	var grains2 int
	var hitsrc bool
	for !hitsrc {
		grains2++
		grid2, hitsrc = dropSand(sandsrc, dim2, grid2)
	}

	for _, row := range grid2 {
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

	// formations = append(formations, []loc{{500, 0}})

	return formations
}

func buildGrid(formations [][]loc, ss loc, offset int) ([]int, [][]rune) {
	// Find grid edges
	minx, maxx := ss.x, ss.x
	miny, maxy := ss.y, ss.y
	for _, f := range formations {
		for _, p := range f {
			if p.x > maxx {
				maxx = p.x
			}

			if p.x < minx {
				minx = p.x
			}

			if p.y > maxy {
				maxy = p.y
			}

			if p.y < miny {
				miny = p.y
			}
		}
	}

	// add l/r margin for easy checking
	minx -= offset
	maxx += offset
	dim := []int{minx, maxx, miny, maxy}

	// Init an empty grid
	var grid [][]rune
	for i := 0; i <= maxy-miny; i++ {
		var col []rune
		for j := 0; j <= maxx-minx; j++ {
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
				grid[p1.y-miny][p1.x-minx] = '#'
				// update a
				step := loc{step(dx), step(dy)}
				p1 = p1.transform(step)
			}
			grid[p2.y-miny][p2.x-minx] = '#'
		}
	}

	// Add sand source
	grid[ss.y-miny][ss.x-minx] = '+'

	return dim, grid
}

func dropSand(ss loc, dim []int, grid [][]rune) ([][]rune, bool) {
	ss = loc{ss.x - dim[0], ss.y - dim[2]}
	grain := ss

	for {
		down := grain.transform(loc{0, 1})
		left := grain.transform(loc{-1, 1})
		right := grain.transform(loc{1, 1})

		//Grain falls in void
		if down.y > dim[3] {
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
			//Grain hits sand source
			if grain == ss {
				return grid, true
			}
			grid[grain.y][grain.x] = 'o'
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
