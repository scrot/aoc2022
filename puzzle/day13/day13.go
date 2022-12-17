package day13

import (
	"bufio"
	"fmt"
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

	input := parseInput(buf)

	var index int
	var correct int

	for i := 0; i < len(input)-1; i += 2 {
		index++
		elemsA := NewList(input[i])
		elemsB := NewList(input[i+1])
		eq := elemsA.CompareTo(elemsB)
		if eq <= 0 {
			correct += index
		}
		fmt.Printf("%s vs %s (compare %d)\n", elemsA, elemsB, eq)
	}

	dpkg := []string{"[[2]]", "[[6]]"}
	input = append(input, dpkg...)

	sort.Slice(input, func(i, j int) bool {
		il := NewList(input[i])
		jl := NewList(input[j])
		return il.LessThan(jl)
	})

	var dp1, dp2 int
	for i, line := range input {
		if line == dpkg[0] {
			dp1 = i + 1
		}
		if line == dpkg[1] {
			dp2 = i + 1
		}
	}

	log.Printf("Answer part I: %d", correct)
	log.Printf("Answer part II: %d", dp1*dp2)
}

func parseInput(input *bufio.Scanner) []string {
	var ls []string
	for input.Scan() {
		l := input.Text()
		if l == "" {
			continue
		}
		ls = append(ls, l)
	}
	return ls
}

func splitList(l []byte) list {
	// fmt.Printf("Split list: %s\n", l)
	var elems list

	var depth int
	var digit []byte
	var ls []byte

	for i := 1; i < len(l)-1; i++ {

		// fmt.Printf("digit: %c, ls: %s\n", digit, string(ls))
		switch l[i] {
		case '[':
			depth++
			ls = append(ls, l[i])
		case ']':
			depth--
			ls = append(ls, l[i])
			if depth == 0 {
				elems = append(elems, ls)
				ls = []byte{}
			}
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if depth == 0 {
				digit = append(digit, l[i])
				if l[i+1] < '0' || l[i+1] > '9' {
					elems = append(elems, digit)
					digit = []byte{}
				}
			} else {
				ls = append(ls, l[i])
			}
		case ',':
			if depth > 0 {
				ls = append(ls, l[i])
			}
		default:
			fmt.Println("Not all chars parsed")
		}
	}

	// fmt.Printf("Result list: %s\n", elems)
	return elems
}

type list [][]byte

func NewList(l string) list {
	return list{[]byte(l)}
}

func (i list) LessThan(j list) bool {
	return i.CompareTo(j) < 0
}

func (i list) EqualTo(j list) bool {
	return i.CompareTo(j) == 0
}

func (i list) CompareTo(j list) int {
	for index := 0; index < len(i); index++ {
		// i is equal to j but longer
		if index >= len(j) {
			return 1
		}

		var res int
		di, erri := strconv.Atoi(string(i[index]))
		dj, errj := strconv.Atoi(string(j[index]))

		// fmt.Printf("Comparing %s and %s\n", i[index], j[index])

		switch {
		case erri == nil && errj != nil: //dj is list
			li := list{i[index]}
			lj := splitList(j[index])
			res = li.CompareTo(lj)
		case erri != nil && errj == nil: //di is list
			li := splitList(i[index])
			lj := list{j[index]}
			res = li.CompareTo(lj)
		case erri != nil && errj != nil: //both are lists
			li := splitList(i[index])
			lj := splitList(j[index])
			res = li.CompareTo(lj)
		default: //both are numbers
			res = cmp(di, dj)
		}

		// if elems not equal return
		if res == 0 {
			continue
		} else {
			return res
		}
	}

	// i is equal to j
	if len(i) == len(j) {
		return 0
	}

	// i is equal but shorter than j
	return -1
}

func cmp(x, y int) int {
	switch {
	case x-y < 0:
		return -1
	case x-y > 0:
		return 1
	default:
		return 0
	}
}
