package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile(
		"test",
		os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	n, err := f.Write([]byte("hello world"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d bytes succeessfully written", n)
}
