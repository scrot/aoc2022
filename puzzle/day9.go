package puzzle

import (
	"bufio"
	"fmt"
	"log"
	"math"
)

type Day9 struct {
	*Day
}

type loc struct {
	x, y int
}

func (h loc) diff(t loc) loc {
	hx, hy := float64(h.x), float64(h.y)
	tx, ty := float64(t.x), float64(t.y)
	return loc{int(math.Abs(hx - tx)), int(math.Abs(hy - ty))}
}

func (l loc) String() string {
	return fmt.Sprintf("(x:%d, y:%d)", l.x, l.y)
}

type vmap map[loc]bool

func (d Day9) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var h, t loc
	v := vmap{loc{0, 0}: true}

	for buf.Scan() {
		move := buf.Text()

		h, t, v = step(h, t, v, move)

		// log.Printf(move)
		// log.Printf("tail %s", t)
		// log.Printf("head %s", h)
	}

	log.Printf("Answer part I: %d", len(v))
	log.Printf("Answer part II: %d", 0)
}

func step(head, tail loc, visited vmap, move string) (loc, loc, vmap) {
	var direction rune
	var steps int
	fmt.Sscanf(move, "%c %d", &direction, &steps)

	for step := 0; step < steps; step++ {
		head = update(head, direction)

		diff := head.diff(tail)
		switch {
		case diff.x <= 1 && diff.y <= 1:
			// Do nothing
		case diff.x == 1 && diff.y == 2:
			tail.x = head.x
			tail = update(tail, direction)
			visited[tail] = true
		case diff.x == 2 && diff.y == 1:
			tail.y = head.y
			tail = update(tail, direction)
			visited[tail] = true
		default:
			tail = update(tail, direction)
			visited[tail] = true
		}
	}
	return head, tail, visited
}

func update(l loc, d rune) loc {
	switch d {
	case 'U':
		l.y++
	case 'D':
		l.y--
	case 'L':
		l.x--
	case 'R':
		l.x++
	}
	return l
}
