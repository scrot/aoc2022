package day6

import (
	"io"
	"log"
	"math/bits"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

func (d Day) Solve() {
	buf, _ := io.ReadAll(d.Dataset)
	defer d.Dataset.Close()

	pkg := markerAt(&buf, 4)
	msg := markerAt(&buf, 14)

	log.Printf("Answer part I: %d\n", pkg)
	log.Printf("Answer part II: %d\n", msg)

}

func markerAt(buf *[]byte, l int) int {
	for i := 0; i < len(*buf)-l-1; i++ {
		// Must be as may bits as alphabet
		var frame uint32

		// Shift bit left char - 'a' positions
		// if double char the OR operation flips bit
		for j := 0; j < l; j++ {
			frame |= 1 << ((*buf)[i+j] - 'a')
		}

		// log.Printf("%b \n", frame)

		// Count unique characters in frame
		if bits.OnesCount32(frame) == l {
			return i + l
		}
	}
	return 0
}
