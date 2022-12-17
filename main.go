package main

import (
	"flag"
	"log"
	"os"

	"github.com/scrot/aoc2022/puzzle"
	"github.com/scrot/aoc2022/puzzle/day1"
	"github.com/scrot/aoc2022/puzzle/day10"
	"github.com/scrot/aoc2022/puzzle/day11"
	"github.com/scrot/aoc2022/puzzle/day12"
	"github.com/scrot/aoc2022/puzzle/day13"
	"github.com/scrot/aoc2022/puzzle/day2"
	"github.com/scrot/aoc2022/puzzle/day3"
	"github.com/scrot/aoc2022/puzzle/day4"
	"github.com/scrot/aoc2022/puzzle/day5"
	"github.com/scrot/aoc2022/puzzle/day6"
	"github.com/scrot/aoc2022/puzzle/day7"
	"github.com/scrot/aoc2022/puzzle/day8"
	"github.com/scrot/aoc2022/puzzle/day9"
)

var (
	token = flag.String("session", "", "Advent of code session token")
	day   = flag.Int("day", 1, "Day of the advent calender to solve")
)

var (
	puzzles = map[int]puzzle.Solver{
		1:  day1.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/1/input")},
		2:  day2.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/2/input")},
		3:  day3.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/3/input")},
		4:  day4.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/4/input")},
		5:  day5.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/5/input")},
		6:  day6.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/6/input")},
		7:  day7.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/7/input")},
		8:  day8.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/8/input")},
		9:  day9.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/9/input")},
		10: day10.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/10/input")},
		11: day11.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/11/input")},
		12: day12.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/12/input")},
		13: day13.Day{Day: puzzle.NewDay("https://adventofcode.com/2022/day/13/input")},
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
