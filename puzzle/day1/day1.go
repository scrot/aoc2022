package day1

import (
	"bufio"
	"log"
	"sort"
	"strconv"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

func (d Day) Solve() {
	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var curCal int
	var topCals []int
	for buf.Scan() {
		l := buf.Text()

		switch len(l) {
		case 0:
			topCals = insertCal(topCals, curCal)
			log.Printf("Elf with %dcal, new top %v\n", curCal, topCals)
			curCal = 0
		default:
			cal, _ := strconv.Atoi(string(l))
			curCal += cal
		}
	}

	totalCals := topCals[0] + topCals[1] + topCals[2]
	log.Printf("Answer Part I: %d\nAnswer Part II: %d\n", topCals[2], totalCals)

}

func insertCal(topCals []int, curCal int) []int {
	if len(topCals) < 3 {
		return append(topCals, curCal)
	}

	sort.Ints(topCals)
	if topCals[0] < curCal {
		topCals[0] = curCal
		return topCals
	}

	return topCals
}
