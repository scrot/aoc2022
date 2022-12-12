package puzzle

import (
	"fmt"
	"io"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day11 struct {
	*Day
}

type monkey struct {
	items     []int
	operation func(int) int
	test      func(int) int

	inspected int
	divider   int
}

func (d Day11) Solve() {
	buf, _ := io.ReadAll(d.Dataset)
	defer d.Dataset.Close()

	monkeys := parseMonkeys(string(buf))
	monkeys2 := parseMonkeys(string(buf))

	log.Printf("Answer part I: %v", goBananas(monkeys, 20, false))
	log.Printf("Answer part II: %d", goBananas(monkeys2, 10000, true))
}

func goBananas(monkeys []*monkey, rounds int, worried bool) int {
	commonDenominator := 1
	for _, m := range monkeys {
		commonDenominator *= m.divider
	}

	for r := 0; r < rounds; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				//monkey inspects
				inspected := item
				inspected = m.operation(inspected)

				if !worried {
					inspected = int(float64(inspected) / 3.0)
				} else {
					// use common denominator to prevent overflow
					inspected %= commonDenominator
				}

				m.inspected++
				m.items = m.items[1:]

				//monkey tests & throws
				newOwner := m.test(inspected)
				monkeys[newOwner].items = append(monkeys[newOwner].items, inspected)
			}
		}
	}

	var activity []int
	for _, m := range monkeys {
		activity = append(activity, m.inspected)
	}
	sort.Ints(activity)
	return activity[len(activity)-1] * activity[len(activity)-2]
}

func parseMonkeys(input string) []*monkey {
	var monkeys []*monkey
	mss := strings.Split(input, "\n\n")

	for _, ms := range mss {
		m := strings.Split(ms, "\n")

		test, div := parseTest(m[3:])
		monkeys = append(monkeys, &monkey{
			items:     parseItems(m[1]),
			operation: parseOperation(m[2]),
			test:      test,
			divider:   div,
		})

	}
	return monkeys
}

func parseItems(l string) []int {
	rx := regexp.MustCompile("[0-9]+")
	matches := rx.FindAllString(l, -1)

	var items []int
	for _, match := range matches {
		item, _ := strconv.Atoi(match)
		items = append(items, int(item))
	}
	return items
}

func parseOperation(l string) func(int) int {
	var operator rune
	var factor string
	fmt.Sscanf(strings.TrimSpace(l), "Operation: new = old %c %s", &operator, &factor)
	return func(worry int) int {
		var f int
		if factor == "old" {
			f = worry
		} else {
			v, _ := strconv.Atoi(factor)
			f = int(v)
		}

		switch operator {
		case '+':
			worry += f
		case '*':
			worry *= f
		default:
			log.Fatal("invalid operator")
		}
		return worry

	}
}

func parseTest(ls []string) (func(int) int, int) {
	var div, t, f int
	fmt.Sscanf(strings.TrimSpace(ls[0]), "Test: divisible by %d", &div)
	fmt.Sscanf(strings.TrimSpace(ls[1]), "If true: throw to monkey %d", &t)
	fmt.Sscanf(strings.TrimSpace(ls[2]), "If false: throw to monkey %d", &f)
	return func(worry int) int {
		if worry%div == 0 {
			return t
		}
		return f
	}, div
}
