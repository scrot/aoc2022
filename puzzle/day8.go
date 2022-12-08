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

	total, score := iterateInnerTrees(grid)
	total += len(grid)*2 + len(grid[0])*2 - 4

	log.Printf("Answer part I: %d", total)
	log.Printf("Answer part II: %d", score)
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

func iterateInnerTrees(grid [][]int) (int, int) {
	var count int
  var highScore int

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
      nd, n := visibleNorth(grid, row, col)
			sd, s := visibleSouth(grid, row, col)
			wd, w := visibleWest(grid, row, col)
			ed, e :=visibleEast(grid, row, col)

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

func visibleEast(grid [][]int, r, c int) (int, bool) {
  var dist int
	for col := c - 1; col >= 0; col-- {
    dist++
		if grid[r][c] <= grid[r][col] {
			return dist, false
		}
	}
	return dist, true
}

func visibleWest(grid [][]int, r, c int) (int, bool) {
  var dist int
	for col := c + 1; col < len(grid[c]); col++ {
    dist++
		if grid[r][c] <= grid[r][col] {
			return dist, false
		}
	}
	return dist,true
}

func visibleNorth(grid [][]int, r, c int) (int, bool) {
  var dist int
	for row := r - 1; row >= 0; row-- {
    dist++
		if grid[r][c] <= grid[row][c] {
			return dist, false
		}
	}
	return dist, true
}

func visibleSouth(grid [][]int, r, c int) (int, bool) {
  var dist int
	for row := r + 1; row < len(grid); row++ {
    dist++
		if grid[r][c] <= grid[row][c] {
			return dist, false
		}
	}
	return dist, true
}
