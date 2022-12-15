package main

import (
	"flag"
	"log"
	"os"

	"github.com/scrot/aoc2022/puzzle"
)

var (
	token = flag.String("session", "", "Advent of code session token")
	day   = flag.Int("day", 1, "Day of the advent calender to solve")
)

var (
	puzzles = map[int]puzzle.Solver{
		1:  puzzle.Day1{Day: puzzle.NewDay("https://adventofcode.com/2022/day/1/input")},
		2:  puzzle.Day2{Day: puzzle.NewDay("https://adventofcode.com/2022/day/2/input")},
		3:  puzzle.Day3{Day: puzzle.NewDay("https://adventofcode.com/2022/day/3/input")},
		4:  puzzle.Day4{Day: puzzle.NewDay("https://adventofcode.com/2022/day/4/input")},
		5:  puzzle.Day5{Day: puzzle.NewDay("https://adventofcode.com/2022/day/5/input")},
		6:  puzzle.Day6{Day: puzzle.NewDay("https://adventofcode.com/2022/day/6/input")},
		7:  puzzle.Day7{Day: puzzle.NewDay("https://adventofcode.com/2022/day/7/input")},
		8:  puzzle.Day8{Day: puzzle.NewDay("https://adventofcode.com/2022/day/8/input")},
		9:  puzzle.Day9{Day: puzzle.NewDay("https://adventofcode.com/2022/day/9/input")},
		10: puzzle.Day10{Day: puzzle.NewDay("https://adventofcode.com/2022/day/10/input")},
		11: puzzle.Day11{Day: puzzle.NewDay("https://adventofcode.com/2022/day/11/input")},
		12: puzzle.Day12{Day: puzzle.NewDay("https://adventofcode.com/2022/day/12/input")},
	}
)

func main() {
	flag.Parse()

	d, ok := puzzles[*day]
	if !ok {
		log.Fatalf("Day doesnt exist in puzzles map")
		return
	}

	if *token != "" {
		d.FetchDataSetByToken(*token)
	} else {
		d.FetchDataByReader(os.Stdin)
	}
	d.Solve()
}
