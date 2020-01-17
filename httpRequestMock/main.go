package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type roundtripper struct{}

func (r roundtripper) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: 555,
		Body:       ioutil.NopCloser(bytes.NewBufferString(`OK`)),
		Header:     make(http.Header),
	}, nil
}

func newHTTPClient() *http.Client {
	return &http.Client{
		Transport: roundtripper{},
	}
}

func main() {
	req, err := http.NewRequest("GET", "http://goosdgsdfghsdfhsdfhgle.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	cli := newHTTPClient()
	response, err := cli.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.StatusCode) // 555
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b)) // "OK"
}
