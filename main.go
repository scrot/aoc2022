package main

import (
  "flag"
	"net/http"
  "io"
	"net/http/cookiejar"
)

var (
  token = flag.String("session", "", "Advent of Code session token")
)

type Day struct {
  PuzzleInput io.ReadCloser
}

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
		"https://adventofcode.com/2022/day/2/input",
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

  var d Day2
  d.PuzzleInput = resp.Body

  d.Solve()
}
