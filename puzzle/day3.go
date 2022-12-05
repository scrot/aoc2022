package puzzle

import (
	"bufio"
	"log"
	"strings"
)

type Day3 struct {
	Day
}

func (d Day3) Solve() {
	prios := generatePriorities()

	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var sum int
	var groupSum int

	counter := 0
	var group []string

	for buf.Scan() {
		l := buf.Text()

		// Part I
		item := findDuplicateItem(l)
		sum += prios[item]

		// Part II
    counter++
    group = append(group, l)
    
    if counter == 3 {
      item := findBadgeItem(group)
      groupSum += prios[item]
      counter = 0
      group = []string{}
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

func findBadgeItem(g []string) rune {
  bucket := make(map[rune]int)
  for _, sack := range g {
    var added string
    for _, c := range sack {
      if !strings.ContainsRune(added, c) {
        bucket[c]++
        added += string(c)
      }
    }
  }
  
  var match []rune
  for key, value := range bucket {
    if value == 3 {
      match = append(match, key)
    }
  }
	log.Printf("%c appears in %v (%d)\n", match, g, len(g))

	return match[0]
}

func prettyPrintPriorities(prios map[rune]int) {
	for k, v := range prios {
		log.Printf("%c, %d\n", k, v)
	}
	log.Printf("Length: %d\n", len(prios))

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
