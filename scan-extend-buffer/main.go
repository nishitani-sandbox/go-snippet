package main

import (
	"bufio"
	"log"
	"os"
)

const (
	maxSize = 12 // bytes
)

func main() {
	f, err := os.Open("txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	max := 2
L:
	log.Println("loop starts")
	f.Seek(0, 0)
	s := bufio.NewScanner(f)
	if err == bufio.ErrTooLong {
		max = max * 2
	}
	if max > maxSize {
		max = maxSize
	}
	s.Buffer(make([]byte, 2), max)

	for s.Scan() {
		log.Println("scan")
		log.Println(s.Text())
	}
	err = s.Err()
	if err == bufio.ErrTooLong {
		if max == maxSize {
			log.Fatal("reach to max size")
		}
		err = s.Err()
		goto L
	}
	if err != nil {
		log.Fatal(err)
	}
}
