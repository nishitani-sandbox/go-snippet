package main

import (
	"log"
	"syscall"
)

func main() {
	var stat syscall.Stat_t
	if err := syscall.Stat("/home/nishitani/src/github.com/nishitani-sandbox/go-snippet/inode/test.txt", &stat); err != nil {
		log.Fatal(err)
	}
	log.Println(stat.Ino)
}
