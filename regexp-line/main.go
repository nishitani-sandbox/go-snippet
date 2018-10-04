package main

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

var (
	hasPkgPat = regexp.MustCompile("(?m)^ii")
)

func main() {
	f, err := os.Open("./dpkg.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	out, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	if !hasPkgPat.Match(out) {
		log.Fatal(errors.New("not matched"))
	}
}
