package main

import (
	"bufio"
	"log"
)

type Day2 struct {
  Day
}

const (
  rock = iota
  paper
  scissors
)

const (
  lose = 0
  draw = 3
  win = 6
)

func(d Day2) Solve() {
  me := map[byte]string{
    'X': "rock",
    'Y': "paper",
    'Z': "scissors",
  }

  elf := map[byte]string{
    'A': "rock",
    'B': "paper",
    'C': "scissors",
  }


  buf := bufio.NewScanner(d.PuzzleInput)
  defer d.PuzzleInput.Close()

  var totalScore int

  for buf.Scan() {
    l := buf.Text()
    
    elfMove := elf[l[0]]
    meMove := me[l[2]]


    score := gameOutcome(meMove, elfMove) + movePoints(meMove)
    totalScore += score
    log.Printf("My move %s, elf move %s (score %d, total %d)\n", meMove, elfMove, score, totalScore)
  }

  log.Printf("Answer Part I: %d", totalScore)
  
}

func gameOutcome(yourMove, opponentMove string) int {
  switch {
  case yourMove == opponentMove:
    return draw
  case yourMove == "rock" && opponentMove == "scissors":
    return win
  case yourMove == "paper" && opponentMove == "rock":
    return win
  case yourMove == "scissors" && opponentMove == "paper":
    return win
  default:
    return lose
  }
}

func movePoints(move string) int {
  switch move {
  case "rock":
    return 1
  case "paper":
    return 2
  case "scissors":
    return 3
  default:
    panic("Not valid")
  }
}
