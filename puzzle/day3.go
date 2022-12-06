package puzzle

import (
	"bufio"
	"log"
	"strings"
)

type Day3 struct {
	*Day
}

func (d Day3) Solve() {
	prios := generatePriorities()

	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var sum int
	var groupSum int

	counter := 0
	var group [3]string

	for buf.Scan() {
		l := buf.Text()

		// Part I
		item := findDuplicateItem(l)
		sum += prios[item]

		// Part II
		group[counter] = l
		counter++

		if counter == 3 {
			item := findBadgeItem(group)
			groupSum += prios[item]
			group = [3]string{}
			counter = 0
		}
	}

	log.Printf("Answer Part I: %d", sum)
	log.Printf("Answer Part II: %d", groupSum)

}

func findDuplicateItem(l string) rune {
	mid := len(l) / 2
	leftCompartment := l[:mid]
	rightCompartment := l[mid:]
	match := strings.IndexAny(leftCompartment, rightCompartment)
	// log.Printf("%c appears in %s and %s\n", rune(l[match]), leftCompartment, rightCompartment)
	return rune(l[match])
}

func findBadgeItem(group [3]string) rune {
	bucket := make(map[rune]int)
	for _, sack := range group {
    added := make(map[rune]bool)
		for _, c := range sack {
			if !added[c] {
				bucket[c]++
				added[c] = true
			}
		}
	}

	var match rune
	for key, value := range bucket {
		if value == 3 {
			match = key
      break
		}
	}
	// log.Printf("%c appears in %v (%d)\n", match, group, len(group))

	return match
}

func generatePriorities() map[rune]int {
	priorities := make(map[rune]int, 52)

	var index int
	for c := 'a'; c <= 'z'; c++ {
		index++
		priorities[c] = index

	}

	for c := 'A'; c <= 'Z'; c++ {
		index++
		priorities[c] = index
	}
	return priorities
}
