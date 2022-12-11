package puzzle

import (
	"bufio"
	"fmt"
	"log"
)

type Day9 struct {
	*Day
}

type loc struct {
	x, y int
}

func (h loc) diff(t loc) loc {
	return loc{h.x - t.x, h.y - t.y}
}

func (l loc) String() string {
	return fmt.Sprintf("(x:%d, y:%d)", l.x, l.y)
}

func (d Day9) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	v := map[loc]bool{{}: true}
	v2 := map[loc]bool{{}: true}

	var rope []loc
	for i := 0; i < 10; i++ {
		rope = append(rope, loc{})
	}

	for buf.Scan() {
		l := buf.Text()
		steps, direction := parse(l)

		for s := 0; s < steps; s++ {
			rope[0] = update(rope[0], direction)
			for k := 1; k < len(rope); k++ {
				rope[k] = follow(rope[k-1], rope[k])
			}
			v[rope[1]] = true
			v2[rope[9]] = true
		}

	}

	log.Printf("Answer part I: %d", len(v))
	log.Printf("Answer part II: %d", len(v2))
}

func parse(move string) (int, rune) {
	var d rune
	var s int
	fmt.Sscanf(move, "%c %d", &d, &s)
	return s, d
}

func follow(head, tail loc) loc {
	d := head.diff(tail)

	if d.x >= -1 && d.x <= 1 && d.y >= -1 && d.y <= 1 {
		return tail
	}

	if d.x > 0 {
		tail = update(tail, 'R')
	} else if d.x < 0 {
		tail = update(tail, 'L')
	}

	if d.y > 0 {
		tail = update(tail, 'U')
	} else if d.y < 0 {
		tail = update(tail, 'D')
	}

	return tail
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
