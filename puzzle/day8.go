package puzzle

import (
	"io"
	"log"
	"strconv"
	"strings"
)

type Day8 struct {
	*Day
}

func (d Day8) Solve() {
	buf, _ := io.ReadAll(d.Dataset)
	defer d.Dataset.Close()

	grid := newGrid(string(buf))

	total := iterateInnerTrees(grid)
	total += len(grid)*2 + len(grid[0])*2 - 4

	log.Printf("Answer part I: %d", total)
	log.Printf("Answer part II: %d", 0)
}

func newGrid(buf string) [][]int {
	lines := strings.Split(string(buf), "\n")
	grid := make([][]int, len(lines)-1)

	for r := 0; r < len(lines); r++ {
		for c := 0; c < len(lines[r]); c++ {
			f, _ := strconv.Atoi(string(lines[r][c]))
			grid[r] = append(grid[r], f)
		}
	}

	return grid
}

func iterateInnerTrees(grid [][]int) int {
	var count int

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			if visibleNorth(grid, row, col) ||
				visibleSouth(grid, row, col) ||
				visibleWest(grid, row, col) ||
				visibleEast(grid, row, col) {
				// log.Printf("Tree (%d, %d) with height %d is visible", row, col, grid[row][col])
				count++
			}
		}
	}

	return count
}

func visibleEast(grid [][]int, r, c int) bool {
	for col := c - 1; col >= 0; col-- {
		if grid[r][c] <= grid[r][col] {
			return false
		}
	}
	return true
}

func visibleWest(grid [][]int, r, c int) bool {
	for col := c + 1; col < len(grid[c]); col++ {
		if grid[r][c] <= grid[r][col] {
			return false
		}
	}
	return true
}

func visibleNorth(grid [][]int, r, c int) bool {
	for row := r - 1; row >= 0; row-- {
		if grid[r][c] <= grid[row][c] {
			return false
		}
	}
	return true
}

func visibleSouth(grid [][]int, r, c int) bool {
	for row := r + 1; row < len(grid); row++ {
		if grid[r][c] <= grid[row][c] {
			return false
		}
	}
	return true
}
