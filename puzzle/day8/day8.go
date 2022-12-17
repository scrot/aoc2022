package day8

import (
	"bufio"
	"log"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

func (d Day) Solve() {
	s := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var buf []string
	for s.Scan() {
		buf = append(buf, s.Text())
	}

	grid := newGrid(buf)

	total, score := treeWalker(grid)

	log.Printf("Answer part I: %d", total)
	log.Printf("Answer part II: %d", score)
}

func newGrid(buf []string) [][]int {
	grid := make([][]int, len(buf))

	for r := 0; r < len(buf); r++ {
		for c := 0; c < len(buf[r]); c++ {
			f := int(buf[r][c])
			grid[r] = append(grid[r], f)
		}
	}

	return grid
}

func treeWalker(grid [][]int) (int, int) {
	var count int
	var highScore int

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			nd, n := checkDirection(grid, row, col, 0, -1)
			sd, s := checkDirection(grid, row, col, 0, 1)
			wd, w := checkDirection(grid, row, col, -1, 0)
			ed, e := checkDirection(grid, row, col, 1, 0)

			score := nd * sd * wd * ed
			if score > highScore {
				highScore = score
			}

			if n || s || w || e {
				count++
			}
		}
	}

	return count, highScore
}

func checkDirection(grid [][]int, r, c, h, v int) (int, bool) {
	var dist int
	for col, row := c+h, r+v; col >= 0 && col < len(grid[c]) && row >= 0 && row < len(grid); col, row = col+h, row+v {
		dist++
		if grid[r][c] <= grid[row][col] {
			return dist, false
		}
	}
	return dist, true
}
