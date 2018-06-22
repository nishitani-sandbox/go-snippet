package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	u, err := url.Parse("https://af.cybozu.com/ubuntu/dists/xenial/InRelease")
	if err != nil {
		log.Fatal(err)
	}
	req := &http.Request{
		Method:     "GET",
		URL:        u,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}
	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.StatusCode)
}
