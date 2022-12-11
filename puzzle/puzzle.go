package puzzle

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
)

const (
	domain = "https://adventofcode.com"
	ghAuth = "https://adventofcode.com/auth/github"
)

type Solver interface {
	Solve()
	FetchDataSetByToken(string)
	FetchDataByReader(io.ReadCloser)
}

type Day struct {
	GHAuth  string
	Url     string
	Dataset io.ReadCloser
}

func NewDay(url string) *Day {
	return &Day{
		Url:    url,
		GHAuth: ghAuth,
	}
}

func (d *Day) FetchDataByReader(input io.ReadCloser) {
	d.Dataset = input
}

func (d *Day) FetchDataSetByToken(sessionToken string) {
	if sessionToken == "" {
		panic(fmt.Errorf("Empty session token"))
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Jar: jar,
	}

	req, err := http.NewRequest(http.MethodGet, d.Url, nil)
	if err != nil {
		panic(err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionToken,
	})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	d.Dataset = resp.Body
}
