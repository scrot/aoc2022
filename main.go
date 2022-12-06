package main

import (
	"flag"

	"github.com/scrot/aoc2022/puzzle"
)

var (
  token = flag.String("session", "", "Advent of Code session token")
)

var (
  puzzles = map[int]puzzle.Solver {
    1: puzzle.Day1{Day: puzzle.NewDay("https://adventofcode.com/2022/day/1/input")},
    2: puzzle.Day2{Day: puzzle.NewDay("https://adventofcode.com/2022/day/2/input")},
    3: puzzle.Day3{Day: puzzle.NewDay("https://adventofcode.com/2022/day/3/input")},
    4: puzzle.Day4{Day: puzzle.NewDay("https://adventofcode.com/2022/day/4/input")},
    5: puzzle.Day5{Day: puzzle.NewDay("https://adventofcode.com/2022/day/5/input")},
    6: puzzle.Day6{Day: puzzle.NewDay("https://adventofcode.com/2022/day/6/input")},
  }
)

func main() {
  flag.Parse()
  
  d := puzzles[5] 
  d.FetchDataSetByToken(*token)
  d.Solve()
}
