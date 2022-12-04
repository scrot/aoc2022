package main

import (
  "flag"
	"net/http"
	"net/http/cookiejar"
)

var (
  token = flag.String("session", "", "Advent of Code session token")
)

type Solver interface {
  Solve()
}

func main() {
  flag.Parse()
  
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Jar: jar,
	}

	session := &http.Cookie{
		Name:  "session",
    Value: *token,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		"https://adventofcode.com/2022/day/1/input",
		nil,
	)

	if err != nil {
		panic(err)
	}

	req.AddCookie(session)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

  d := Day1{
    PuzzleInput: resp.Body,
  }

  d.Solve()
}
