package puzzle

import (
	"bufio"
	"fmt"
	"log"
	"sort"
)

type Day12 struct {
	*Day
}

type place struct {
	index  int
	height int
}

func (p place) String() string {
	return fmt.Sprintf("%d", p.index)
}

func (d Day12) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	// input > start, end, indices
	start, end, grid := parseInput(buf)

	// indices > adjacency list
	adjacencyList := adjacencyList(grid)

	// all squares with height 'a'
	var lowestPlaces []place
	for _, row := range grid {
		for _, field := range row {
			if field.height == 0 {
				lowestPlaces = append(lowestPlaces, field)
			}
		}
	}

	var p1 int
	var ps []int
	for _, p := range lowestPlaces {
		if p == start {
			p1, _ = bfs(p, end, adjacencyList)
		}

		path, _ := bfs(p, end, adjacencyList)
		if path != -1 {
			ps = append(ps, path)
		}
	}

	sort.Ints(ps)

	log.Printf("Answer part I: %d", p1)
	log.Printf("Answer part II: %d", ps[0])

}

func parseInput(input *bufio.Scanner) (place, place, [][]place) {
	var index int
	var row int
	var grid [][]place
	var start, end place

	for input.Scan() {
		line := input.Text()

		grid = append(grid, make([]place, 0, len(line)))
		for _, char := range line {
			p := place{
				index:  index,
				height: charToHeight(char),
			}

			if char == 'S' {
				start = p
			}

			if char == 'E' {
				end = p
			}

			grid[row] = append(grid[row], p)
			index++
		}
		row++
	}

	return start, end, grid
}

func adjacencyList(grid [][]place) map[int][]place {
	// indices > adjacency list
	adjacencyList := make(map[int][]place)
	for r, row := range grid {
		for c, cp := range row {
			// iterate over all neighbor places
			var neighbors []place

			dr := []int{1, -1, 0, 0}
			dc := []int{0, 0, 1, -1}

			for d := 0; d < 4; d++ {
				nr := r + dr[d]
				nc := c + dc[d]

				// only if within grid boundaries
				if nr >= 0 && nc >= 0 && nr < len(grid) && nc < len(grid[r]) {
					np := grid[nr][nc]
					dh := np.height - cp.height

					// only if at most 1 higher
					if dh <= 1 {
						neighbors = append(neighbors, np)
					}
				}
			}

			// add node lo adjacency list
			adjacencyList[cp.index] = neighbors
		}
	}
	return adjacencyList
}

func bfs(start, end place, adjacencyList map[int][]place) (int, []place) {
	var queue []place
	queue = append(queue, start)

	visited := make(map[int]bool)
	visited[start.index] = true

	previous := make(map[int]place)

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		//find shortest path if goal reached
		if current == end {
			var shortest []place
			for current != start {
				shortest = append(shortest, current)
				current = previous[current.index]
				// fmt.Println(shortest)
			}
			return len(shortest), shortest
		}

		// enqueue neighbors of current node
		for _, neighbor := range adjacencyList[current.index] {
			if !visited[neighbor.index] {
				previous[neighbor.index] = current
				queue = append(queue, neighbor)
				visited[neighbor.index] = true
			}
		}

	}
	// fmt.Println("Empty queue, end not found")
	return -1, nil
}

func charToHeight(char rune) int {
	if char == 'S' {
		char = 'a'
	}

	if char == 'E' {
		char = 'z'
	}

	return int(char) - int('a')
}
