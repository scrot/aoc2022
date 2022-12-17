package day2

import (
	"bufio"
	"log"

	"github.com/scrot/aoc2022/puzzle"
)

type Day struct {
	*puzzle.Day
}

const (
	lose = 0
	draw = 3
	win  = 6
)

func (d Day) Solve() {
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

	res := map[byte]int{
		'X': lose,
		'Y': draw,
		'Z': win,
	}

	buf := bufio.NewScanner(d.Dataset)
	defer d.Dataset.Close()

	var totalScore int
	var totalScoreAlt int

	for buf.Scan() {
		l := buf.Text()

		elfMove := elf[l[0]]
		meMove := me[l[2]]
		score := gameOutcome(meMove, elfMove) + movePoints(meMove)
		totalScore += score
		// log.Printf("My move %s, elf move %s (score %d, total %d)\n", meMove, elfMove, score, totalScore)

		outcome := res[l[2]]
		meMoveAlt := inferMove(elfMove, outcome)
		scoreAlt := gameOutcome(meMoveAlt, elfMove) + movePoints(meMoveAlt)
		totalScoreAlt += scoreAlt
		// log.Printf("Elf move %s, required outcome %d, my move %s (score %d, total %d)\n", elfMove, outcome, meMoveAlt, scoreAlt, totalScoreAlt)

	}

	log.Printf("Answer Part II: %d", totalScoreAlt)

}

func inferMove(opponentMove string, outcome int) string {
	switch {
	case outcome == draw:
		return opponentMove
	case opponentMove == "rock" && outcome == win:
		return "paper"
	case opponentMove == "paper" && outcome == win:
		return "scissors"
	case opponentMove == "scissors" && outcome == win:
		return "rock"
	case opponentMove == "rock" && outcome == lose:
		return "scissors"
	case opponentMove == "paper" && outcome == lose:
		return "rock"
	case opponentMove == "scissors" && outcome == lose:
		return "paper"
	default:
		return ""
	}
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
