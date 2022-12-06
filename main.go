package main

import (
	"flag"

	"github.com/scrot/aoc2022/puzzle"
)

var (
  token = flag.String("session", "", "Advent of Code session token")
)

var (
  day1 = puzzle.Day1{Day: puzzle.NewDay("https://adventofcode.com/2022/day/1/input")}
  day2 = puzzle.Day2{Day: puzzle.NewDay("https://adventofcode.com/2022/day/2/input")}
  day3 = puzzle.Day3{Day: puzzle.NewDay("https://adventofcode.com/2022/day/3/input")}
  day4 = puzzle.Day4{Day: puzzle.NewDay("https://adventofcode.com/2022/day/4/input")}
)

func main() {
  flag.Parse()
  
  d := day4 
  d.FetchDataSetByToken(*token)
  d.Solve()
}
