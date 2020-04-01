package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var s strings.Builder
	n, err := io.Copy(&s, f)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d: %s", n, s.String())
	log.Printf("%d: %s", n, s.String())
}
