package puzzle

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

const (
	domain = "https://adventofcode.com"
	ghAuth = "https://adventofcode.com/auth/github"
)

type Solver interface {
	Solve()
}

type Day struct {
	GHAuth  string
	Url     string
	Dataset io.ReadCloser
}

func NewDay(url string) Day {
	return Day{
		Url:    url,
		GHAuth: ghAuth,
	}
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
    Name: "session",
    Value: sessionToken,
  })

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

  d.Dataset = resp.Body
}

func (d *Day) FetchDataSet() {
	// Client with cookiejar
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Jar: jar,
	}

	resp, err := client.Get(d.GHAuth)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	url, err := url.Parse(domain)
	if err != nil {
		panic(err)
	}

	client.Jar.SetCookies(url, resp.Cookies())
	for _, cookie := range client.Jar.Cookies(url) {
		log.Printf("Cookie %s: %s\n\n", cookie.Name, cookie.Value)
	}

	for _, header := range resp.Header {
		log.Printf("Header: %s\n\n", header)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n\n", body)

}
