package puzzle

import (
	"bufio"
	"log"
	"strconv"
	"strings"
)

type Day4 struct {
	*Day
}

func (d Day4) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var contains, overlaps int

	for buf.Scan() {
		l := buf.Text()

		sections := strings.Split(l, ",")
		s1 := strings.Split(sections[0], "-")
		s2 := strings.Split(sections[1], "-")

		s1s, _ := strconv.Atoi(s1[0])
		s1e, _ := strconv.Atoi(s1[1])

		s2s, _ := strconv.Atoi(s2[0])
		s2e, _ := strconv.Atoi(s2[1])

		if sectionContains(s1s, s1e, s2s, s2e) {
			contains++
		}

    if sectionOverlaps(s1s, s1e, s2s, s2e) {
      overlaps++
    }
	}

	log.Printf("Answer Part I: %d", contains)
	log.Printf("Answer Part II: %d", overlaps)

}

func sectionContains(s1s, s1e, s2s, s2e int) bool {
	switch {
	case s1s >= s2s && s1e <= s2e:
		return true
	case s2s >= s1s && s2e <= s1e:
		return true
	default:
		return false
	}
}

func sectionOverlaps(s1s, s1e, s2s, s2e int) bool {
	switch {
	case s1e < s2s:
		return false
	case s2e < s1s:
		return false
	default:
		return true
	}
}
