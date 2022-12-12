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
}

func (d Day11) Solve() {
	buf, _ := io.ReadAll(d.Dataset)
	defer d.Dataset.Close()

	monkeys := parseMonkeys(string(buf))
	for _, m := range monkeys {
		log.Println(m.items)
	}

	for r := 0; r < 20; r++ {
		for _, m := range monkeys {
			for _, item := range m.items {
				//monkey inspects
				inspected := item
				inspected = m.operation(inspected)
				m.inspected++
				m.items = m.items[1:]

				//monkey tests & throws
				newOwner := m.test(inspected)
				monkeys[newOwner].items = append(monkeys[newOwner].items, inspected)
				log.Printf("Monkey %d items: %v", newOwner, monkeys[newOwner].items)
			}
		}
	}

	var activity []int
	for _, m := range monkeys {
		activity = append(activity, m.inspected)
	}
	sort.Ints(activity)

	log.Printf("Answer part I: %v", activity[len(activity)-1]*activity[len(activity)-2])
	log.Printf("Answer part II: %d", 0)
}

func parseMonkeys(input string) []*monkey {
	var monkeys []*monkey
	mss := strings.Split(input, "\n\n")

	for _, ms := range mss {
		m := strings.Split(ms, "\n")

		monkeys = append(monkeys, &monkey{
			items:     parseItems(m[1]),
			operation: parseOperation(m[2]),
			test:      parseTest(m[3:]),
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
		items = append(items, item)
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
			f, _ = strconv.Atoi(factor)
		}
		switch operator {
		case '+':
			worry += f
		case '*':
			worry *= f
		default:
			log.Fatal("invalid operator")
		}
		return int(float64(worry) / 3.0)

	}
}

func parseTest(ls []string) func(int) int {
	var div, t, f int
	fmt.Sscanf(strings.TrimSpace(ls[0]), "Test: divisible by %d", &div)
	fmt.Sscanf(strings.TrimSpace(ls[1]), "If true: throw to monkey %d", &t)
	fmt.Sscanf(strings.TrimSpace(ls[2]), "If false: throw to monkey %d", &f)
	return func(worry int) int {

		if worry%div == 0 {
			return t
		}
		return f
	}
}
