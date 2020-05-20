package main

import (
	"log"
	"syscall"
)

func main() {
	mem, err := syscall.Mmap(-1, 0, 1<<20, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		log.Fatal(err)
	}
	syscall.Munmap(mem)
}
