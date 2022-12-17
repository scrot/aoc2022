package day13

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

	input := parseInput(buf)

	var index, correct int
	for i := 0; i < len(input)-1; i += 2 {
		index++
		elemsA := [][]byte{input[i]}
		elemsB := [][]byte{input[i+1]}
		if compare(elemsA, elemsB) <= 0 {
			correct += index
		}
	}

	input = append(input, []byte("[[2]]"), []byte("[[6]]"))

	sort.Slice(input, func(i, j int) bool {
		il := [][]byte{input[i]}
		jl := [][]byte{input[j]}
		return compare(il, jl) <= 0
	})

	var dp1, dp2 int
	for i, line := range input {
		if string(line) == "[[2]]" {
			dp1 = i + 1
		}
		if string(line) == "[[6]]" {
			dp2 = i + 1
		}
	}

	log.Printf("Answer part I: %d", correct)
	log.Printf("Answer part II: %d", dp1*dp2)
}

func parseInput(input *bufio.Scanner) [][]byte {
	var ls [][]byte
	for input.Scan() {
		l := input.Text()
		if l == "" {
			continue
		}
		ls = append(ls, []byte(l))
	}
	return ls
}

func splitList(l []byte) [][]byte {
	var elems [][]byte

	var depth int
	var digit, ls []byte

	for i := 1; i < len(l)-1; i++ {

		if l[i] == '[' {
			depth++
		}

		if depth > 0 {
			ls = append(ls, l[i])
		}

		if l[i] == ']' {
			depth--

			if depth == 0 {
				elems = append(elems, ls)
				ls = []byte{}
			}
		}

		if l[i] >= '0' && l[i] <= '9' && depth == 0 {
			digit = append(digit, l[i])
			if l[i+1] < '0' || l[i+1] > '9' {
				elems = append(elems, digit)
				digit = []byte{}
			}
		}
	}

	return elems
}

func compare(i, j [][]byte) int {
	for index := 0; index < len(i) && index < len(j); index++ {
		di, erri := strconv.Atoi(string(i[index]))
		dj, errj := strconv.Atoi(string(j[index]))

		var res int
		var li, lj [][]byte
		switch {
		case erri == nil && errj != nil:
			li = [][]byte{i[index]}
			lj = splitList(j[index])
			res = compare(li, lj)
		case erri != nil && errj == nil:
			li = splitList(i[index])
			lj = [][]byte{j[index]}
			res = compare(li, lj)
		case erri != nil && errj != nil:
			li = splitList(i[index])
			lj = splitList(j[index])
			res = compare(li, lj)
		default:
			res = di - dj
		}

		if res != 0 {
			return res
		}
	}

	return len(i) - len(j)

}
