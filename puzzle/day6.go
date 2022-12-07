package puzzle

import (
	"bufio"
	"log"
)

type Day6 struct {
	*Day
}

func (d Day6) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()
	buf.Split(bufio.ScanRunes)

	var sFrame, mFrame string
	var foundS, foundM bool
	var index int

	for buf.Scan() {
		r := buf.Text()
		index++
		// log.Printf("%d: %s", index, r)

		if len(sFrame) < 4 {
			sFrame += r
		} else {
			if !foundS {
				sFrame = sFrame[1:] + r
				if !containsDuplicates(sFrame) {
					log.Printf("Answer part I: %d (%s)\n", index, sFrame)
					foundS = true
				}
			}
		}

		if len(mFrame) < 14 {
			mFrame += r
		} else {
			if !foundM {
				mFrame = mFrame[1:] + r
				if !containsDuplicates(mFrame) {
					log.Printf("Answer part II: %d (%s)\n", index, mFrame)
					foundM = true
				}
			}
		}
	}
}

func containsDuplicates(frame string) bool {
	for i, r := range frame {
		for j := i + 1; j < len(frame); j++ {
			if rune(frame[j]) == r {
				return true
			}
		}
	}
	return false
}
