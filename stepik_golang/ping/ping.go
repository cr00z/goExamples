package main

import (
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Head(url string) (resp *http.Response, err error)
}

type Pinger struct {
	client HTTPClient
}

func (p Pinger) Ping(url string) bool {
	resp, err := p.client.Head(url)
	if err != nil {
		return false
	}
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

func main() {
	client := &http.Client{}
	pinger := Pinger{client}
	url := "https://ya.ru"
	alive := pinger.Ping(url)
	fmt.Println(url, "is alive =", alive)
}